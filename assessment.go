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

var i int
var metering_point_id int
var reading_type int
var reading int
var created_at int64
var totalCost float64

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

	r := csv.NewReader(csvFile) //initializes reader

	for {
		i += 1 //basic counter
		fmt.Println(i, "---------------------------")

		//reads one (1) line from data file
		record, err := r.Read()

		//breaks out of loop when EOF has been reached and writes the final reading
		if err == io.EOF {
			csvWriter(totalCost, metering_point_id)
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		if i != 1 { // skips first line (column titles)

			if i > 2 { // checks if at least 2 lines have been read so that they can be compared
				newMeteringPointId, _ := strconv.Atoi(record[0])

				if newMeteringPointId != metering_point_id { //checks if new reading is from a new meter ID. If so, writes total cost to file and resets counter
					csvWriter(totalCost, metering_point_id)
					totalCost = 0
					i = 2
				}

				if reading_type == 1 {
					electricity_1 = reading
				} else if reading_type == 2 {
					gas_1 = reading
				}

				newReading, _ := strconv.Atoi(record[2])
				newReadingType, _ := strconv.Atoi(record[1])

				if newReadingType == 1 {
					electricity_2 := newReading
					totalCost += electricityCost(electricity_2, electricity_1, created_at)
				} else if newReadingType == 2 {
					gas_2 = newReading
					totalCost += gasCost(gas_2, gas_1, created_at)
				}
				fmt.Println("Total cost =", totalCost)

			}

		}

		metering_point_id, _ = strconv.Atoi(record[0])
		reading_type, _ = strconv.Atoi(record[1])
		reading, _ = strconv.Atoi(record[2])
		created_at, _ = strconv.ParseInt(record[3], 10, 64)

	}
}

func electricityCost(newReading int, oldReading int, created_at int64) float64 {
	var cost float64

	usage := float64(newReading - oldReading)
	kWh := usage / 1000

	fmt.Println("Electricity usage =", usage, "kWh =", kWh)

	if usage <= 100 && usage >= 0 {
		if weekday(created_at) && rate(created_at) {
			cost = kWh * 0.20
		} else {
			cost = kWh * 0.18
		}
	} else {
		cost = 0
	}

	return cost
}

func gasCost(newReading int, oldReading int, created_at int64) float64 {
	var cost float64

	usage := float64(newReading - oldReading)
	kWh := usage * 9.769

	fmt.Println("Gas usage =", usage, "kWh =", kWh)

	if usage <= 100 && usage >= 0 {
		cost = kWh * 0.06
	}

	return cost
}

func weekday(created_at int64) bool {
	var isWeekday bool
	day := int(time.Unix(created_at, 0).Weekday())

	if day != 0 || day != 6 {
		isWeekday = true
	} else {
		isWeekday = false
	}
	return isWeekday
}

func rate(created_at int64) bool {
	var fullRate bool

	hour, _, _ := time.Unix(created_at, 0).Clock()

	if hour >= 7 && hour <= 23 {
		fullRate = true
	} else {
		fullRate = false
	}
	return fullRate
}

func csvWriter(totalCost float64, metering_point_id int) {
	csvFile, err := os.OpenFile("output.csv", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println(err)
	}
	defer csvFile.Close()

	var id string = strconv.Itoa(metering_point_id)
	var cost string = strconv.FormatFloat(totalCost, 'f', 2, 64)
	fmt.Println(id, cost)

	w := csv.NewWriter(csvFile)
	w.Write([]string{id, cost})
	w.Flush()
}
