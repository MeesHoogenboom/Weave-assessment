package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

// type energyData struct {
// 	metering_point_id string
// 	reading_type      string
// 	reading           string
// 	created_at        string
// }

var i int

func main() {
	csvReader()
}

func csvReader() {
	//opens .csv file found in same folder as .go program
	csvFile, err := os.Open("data.csv")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Successfully opened CSV file")
	}
	//closes file after use
	defer csvFile.Close()

	//initializes reader of .csv file
	r := csv.NewReader(csvFile)

	//infinite for-loop over all records found in file
	for {
		i += 1

		//reads one record, preferable to r.ReadAll() to save memory when processing bigger data-sets
		record, err := r.Read()

		//breaks out of loop when reaching the end of file
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal("error=", err)
		}
		if i != 1 {
			// data := energyData{
			// 	metering_point_id: record[0], //string
			// 	reading_type:      record[1],
			// 	reading:           record[2],
			// 	created_at:        record[3],
			// }

			metering_point_id, _ := strconv.Atoi(record[0])
			reading_type, _ := strconv.Atoi(record[1])
			reading, _ := strconv.Atoi(record[2])
			created_at, _ := strconv.Atoi(record[3])

			fmt.Println(metering_point_id, reading_type, reading, created_at)
		}

	}
}
