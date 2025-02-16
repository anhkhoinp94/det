package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/xuri/excelize/v2"
)

type Data struct {
	Id  int    `json:"id"`
	En1 string `json:"en1"`
	En2 string `json:"en2"`
	En3 string `json:"en3"`
	En4 string `json:"en4"`
	Vn1 string `json:"vn1"`
}

func main() {
	// Open the Excel file
	file, err := excelize.OpenFile("AWL-DETVN.xlsx")
	if err != nil {
		log.Fatalf("Failed to open file: %v", err)
	}
	defer file.Close()

	// Get all rows from the first sheet
	sheetName := file.GetSheetName(0) // Get the first sheet
	rows, err := file.GetRows(sheetName)
	if err != nil {
		log.Fatalf("Failed to get rows: %v", err)
	}

	var data []Data

	// Iterate from line 3 (index 2, since it's zero-based)
	for i := 3; i < len(rows); i++ {
		row := rows[i]
		if len(row) < 3 { // Ensure we have enough columns
			continue
		}
		data = append(data, Data{Id: i - 3, En1: row[0], En2: row[6], En3: "", En4: row[4], Vn1: "(" + row[2] + ")" + " " + row[5]})
	}

	// Convert to JSON
	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		log.Fatalf("Failed to convert to JSON: %v", err)
	}

	// Write JSON to file
	jsonFileName := "output.json"
	if err := os.WriteFile(jsonFileName, jsonData, 0644); err != nil {
		log.Fatalf("Failed to write JSON file: %v", err)
	}

	fmt.Printf("JSON data has been written to %s\n", jsonFileName)
}
