package main

import (
	"encoding/csv"
	"log"
	"os"
)

const (
	baseDir  = "session-16-introduction-read-write-csv"
	fileName = "data.csv"
)

func main() {
	// Read CSV file
	csvFile, err := os.Open("./" + fileName)
	if err != nil {
		log.Fatal("Error opening file: ", err)
	}

	defer csvFile.Close()
	// Write CSV file

	lines, err := csv.NewReader(csvFile).ReadAll()
	if err != nil {
		log.Fatal("Error reading file: ", err)
	}

	for _, line := range lines {
		name, email := line[0], line[1]

		log.Printf("Name: %s, Email: %s \n", name, email)
	}
}
