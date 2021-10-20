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
var totalCost, cost float64

var electricity_1, electricity_2 int
var gas_1, gas_2 int

func main() {
	csvReader()
}

// reads csv file line by line
func csvReader() {
	csvFile, err := os.Open("data.csv")
	if err != nil {
		fmt.Println(err)
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
			csvWriter(totalCost, meteringPointId)
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		// skips first line (column titles)
		if i != 1 {

			// checks if at least 2 lines have been read so that they can be compared
			if i > 2 {
				newMeteringPointId, _ = strconv.Atoi(record[0])

				//checks if new reading is from a new meter ID. If so, writes total cost to file and resets counter
				if newMeteringPointId != meteringPointId {
					csvWriter(totalCost, meteringPointId)
					totalCost = 0
					i = 2
				}

				if readingType == 1 {
					electricity_1 = reading
				} else if readingType == 2 {
					gas_1 = reading
				}

				newReading, _ = strconv.Atoi(record[2])
				newReadingType, _ = strconv.Atoi(record[1])

				if newReadingType == 1 {
					electricity_2 = newReading
					cost, readingSkipped = cost(electricity_2, electricity_1, createdAt, newReadingType)
					totalCost += cost
				} else if newReadingType == 2 {
					gas_2 = newReading
					cost, readingSkipped = cost(gas_2, gas_1, createdAt, newReadingType)
					totalCost += cost
				}

			}

		}

		meteringPointId, _ = strconv.Atoi(record[0])
		readingType, _ = strconv.Atoi(record[1])
		reading, _ = strconv.Atoi(record[2])
		createdAt, _ = strconv.ParseInt(record[3], 10, 64)

	}
}

//returns cost of electricity based on two readings and the time of the (week)day
func electricityCost(usage int, createdAt int64) float64 {
	var cost float64

	kWh := usage / 1000

		//checks which tarif needs to be used
		if weekday(createdAt) && rate(createdAt) {
			cost = kWh * 0.20
		} else {
			cost = kWh * 0.18
		}

	return cost
}

func gasCost(usage int) float64 {
	var cost float64

	kWh := usage * 9.769
	cost = kWh * 0.06

	return cost
}

func cost(newReading int, oldReading int, createdAt int64, readingType, int) float64, bool {
	var cost float64
	var readingSkipped bool
	usage := float64(newReading - oldReading)
	
	
	if usage <= 100 && usage >= 0 && readingSkipped == false {
		if readingType == 1 {

			cost = electricityCost(usage, createdAt)

		} else if readingType == 2 {

			cost = gasCost(usage)

		}
	} else if readingSkipped && usage >= 0 {
		if readingType == 1 {

			cost = electricityCost(usage, createdAt)

		} else if readingType == 2 {
			
			cost = gasCost(usage)
		}
	} else readingSkipped == true

	return cost readingSkipped		
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

//returns true if unix timestamp is during peak hours
func rate(createdAt int64) bool {
	var fullRate bool

	hour, _, _ := time.Unix(createdAt, 0).Clock()

	if hour >= 7 && hour <= 23 {
		fullRate = true
	}
	return fullRate
}

//writes meter ID and total cost associated with that ID to file
func csvWriter(totalCost float64, meteringPointId int) {
	//creates file if none exists, otherwise appends to output file
	csvFile, err := os.OpenFile("output.csv", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println(err)
	}
	defer csvFile.Close()

	var id string = strconv.Itoa(meteringPointId)
	var cost string = strconv.FormatFloat(totalCost, 'f', 2, 64)
	fmt.Println(id, cost)

	w := csv.NewWriter(csvFile)
	w.Write([]string{id, cost})
	w.Flush()
}
