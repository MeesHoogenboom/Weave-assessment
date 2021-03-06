package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"time"
)

var record []string
var meteringPointId, newMeteringPointId int
var readingType, newReadingType int
var reading, newReading int
var createdAt int64
var totalCost, price float64
var electricityReadingSkipped, gasReadingSkipped bool

var electricity_1, electricity_2 int
var gas_1, gas_2 int

func main() {
	csvReader("data.csv")
}

// reads csv file line by line
func csvReader(fileName string) {
	//creates a YYYYMMDDHHMMSS timestamp for the output file
	timestamp := "output_" + (time.Now().Format("20060102150405")) + ".csv"

	csvFile, err := os.Open(fileName)
	fmt.Printf("Opened file %v\n", fileName)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer csvFile.Close()

	//initializes reader
	r := csv.NewReader(csvFile)

	var i int

	for {
		//basic counter
		i += 1

		//reads one (1) line from data file
		record, err = r.Read()

		//breaks out of loop when EOF has been reached and writes the final reading to file
		if err == io.EOF {
			csvWriter(totalCost, meteringPointId, timestamp)
			fmt.Println("Done!")
			break
		}
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}

		// skips first line (column titles)
		if i != 1 {

			// checks if at least 2 lines have been read so that they can be compared
			if i > 2 {
				newMeteringPointId, _ = strconv.Atoi(record[0])

				//checks if new reading is from a new meter ID. If so, writes total cost to file and resets counters
				if newMeteringPointId != meteringPointId {
					csvWriter(totalCost, meteringPointId, timestamp)

					electricityReadingSkipped, gasReadingSkipped = false, false
					totalCost = 0
					i = 2
				}

				newReading, _ = strconv.Atoi(record[2])
				newReadingType, _ = strconv.Atoi(record[1])

				//checks wether the last reading was skipped, if so, retains the previous reading. Otherwise assigns the reading to gas or electricity
				if readingType == 1 && !electricityReadingSkipped {
					electricity_1 = reading
				} else if readingType == 2 && !gasReadingSkipped {
					gas_1 = reading
				} else {
					if readingType == 1 {
						electricity_1 = electricity_2
					} else if readingType == 2 {
						gas_1 = gas_2
					}
				}

				//calculates cost based upon reading type and adds everything together
				if newReadingType == 1 {
					electricity_2 = newReading
					if electricity_1 != 0 {
						price, electricityReadingSkipped = cost(electricity_2, electricity_1, createdAt, newReadingType, electricityReadingSkipped)
						totalCost += price
					}
				} else if newReadingType == 2 {
					gas_2 = newReading
					if gas_1 != 0 {
						price, gasReadingSkipped = cost(gas_2, gas_1, createdAt, newReadingType, gasReadingSkipped)
						totalCost += price
					}
				}

			}

		}
		//set ups current variables as previous ones for the next loop
		meteringPointId, _ = strconv.Atoi(record[0])
		readingType, _ = strconv.Atoi(record[1])
		reading, _ = strconv.Atoi(record[2])
		createdAt, _ = strconv.ParseInt(record[3], 10, 64)

	}
}

//validates usage data and calculates the energy cost based on type
func cost(newReading int, oldReading int, createdAt int64, readingType int, readingSkipped bool) (float64, bool) {
	var price float64
	usage := float64(newReading - oldReading)

	//checks wether usage is a correct reading
	if usage <= 100 && usage >= 0 && !readingSkipped {
		if readingType == 1 {

			price = electricityCost(usage, createdAt)

		} else if readingType == 2 {

			price = gasCost(usage)

		}

		// if last reading was invalid, does away with the criteria that usage has to be less than 100 since the usage now contains 2 time units
	} else if readingSkipped && usage >= 0 {
		if readingType == 1 {

			price = electricityCost(usage, createdAt)

		} else if readingType == 2 {

			price = gasCost(usage)
		}
	}
	if price != float64(0) {
		readingSkipped = false
	}

	return price, readingSkipped
}

//returns cost of electricity based on usage and the time of the (week)day
func electricityCost(usage float64, createdAt int64) float64 {
	var cost float64

	kWh := usage / 1000

	//checks which rate needs to be used
	if weekday(createdAt) && rate(createdAt) {
		cost = kWh * 0.20
	} else {
		cost = kWh * 0.18
	}

	return cost
}

//returns gas cost
func gasCost(usage float64) float64 {
	var cost float64

	kWh := usage * 9.769
	cost = kWh * 0.06

	return cost
}

//returns true if Unix timestamp is a weekday (mo, tu, we, th, fr)
func weekday(createdAt int64) bool {
	var isWeekday bool = true
	day := int(time.Unix(createdAt, 0).Weekday())

	if day == 0 || day == 6 {
		isWeekday = false
	}
	return isWeekday
}

//returns true if unix timestamp is during 'peak' hours
func rate(createdAt int64) bool {
	var fullRate bool

	hour, _, _ := time.Unix(createdAt, 0).Clock()

	if hour >= 7 && hour <= 23 {
		fullRate = true
	}
	return fullRate
}

//writes meter ID and total cost associated with that ID to file
func csvWriter(totalCost float64, meteringPointId int, timestamp string) {
	//creates file if none exists, otherwise appends to output file
	csvFile, err := os.OpenFile(timestamp, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println(err)
	}
	defer csvFile.Close()

	// formats output data
	id := strconv.Itoa(meteringPointId)
	cost := strconv.FormatFloat(totalCost, 'f', 2, 64)

	output := []string{id, cost}

	//initializes writer
	w := csv.NewWriter(csvFile)
	w.Write(output)
	w.Flush()
	fmt.Println(output)
}
