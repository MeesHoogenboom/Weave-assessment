package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

type energy struct {
	meteringPointId int
	type1           int
	reading         int
	createdAt       int
}

func main() {
	csvReader()
}

func csvReader() {
	csvFile, err := os.Open("data.csv")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Succesfully opened CSV file!")
	}
	defer csvFile.Close()

	reader := csv.NewReader(csvFile)
	reader.FieldsPerRecord = 0

	records, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(records)
}
