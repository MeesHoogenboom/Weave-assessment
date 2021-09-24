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

func main() {
	csvReader()
}

func csvReader() {
	csvFile, err := os.Open("data.csv")
	if err != nil {
		fmt.Println(err)
	}

	defer csvFile.Close()

	r := csv.NewReader(csvFile)

	for {
		i += 1

		record, err := r.Read()

		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		if i != 1 {
			if i > 2 {
				newReading, _ := strconv.Atoi(record[2])
				// cost := electricityCost(newReading, reading)
				fmt.Println(newReading)
				weekday(created_at)

			}
			metering_point_id, _ = strconv.Atoi(record[0])
			reading_type, _ = strconv.Atoi(record[1])
			reading, _ = strconv.Atoi(record[2])
			created_at, _ = strconv.ParseInt(record[3], 10, 64)

			fmt.Println("new reading =", reading)
			// fmt.Println(metering_point_id, reading_type, reading, created_at)
		}

	}
}

// func electricityCost(newReading int, oldReading int) {
// 	usage := newReading - oldReading
// 	if usage > 100 || usage < 0 {

// 	}
// }

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
