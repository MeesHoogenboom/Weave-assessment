package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

var i int
var metering_point_id int
var reading_type int
var reading int
var created_at int

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
				new_reading, _ := strconv.Atoi(record[2])
				usage := new_reading - reading
				fmt.Println("usage =", usage)
			}
			metering_point_id, _ = strconv.Atoi(record[0])
			reading_type, _ = strconv.Atoi(record[1])
			reading, _ = strconv.Atoi(record[2])
			created_at, _ = strconv.Atoi(record[3])

			fmt.Println("new reading =", reading)
			// fmt.Println(metering_point_id, reading_type, reading, created_at)
		}

	}
}
