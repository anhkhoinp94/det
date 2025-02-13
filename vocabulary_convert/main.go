package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type Item struct {
	Id  int    `json:"id"`
	En1 string `json:"en1"`
	En2 string `json:"en2"`
	En3 string `json:"en3"`
	En4 string `json:"en4"`
	Vn1 string `json:"vn1"`
}

func main() {
	// Open the file
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	var items []Item
	scanner := bufio.NewScanner(file)
	idCounter := 867

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ": ")
		if len(parts) == 2 {
			item := Item{
				Id:  idCounter,
				En1: parts[0],
				Vn1: parts[1],
			}
			items = append(items, item)
			idCounter++
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Convert to JSON
	jsonData, err := json.MarshalIndent(items, "", "  ")
	if err != nil {
		fmt.Println("Error converting to JSON:", err)
		return
	}

	// Write JSON to file
	jsonFile, err := os.Create("output.json")
	if err != nil {
		fmt.Println("Error creating JSON file:", err)
		return
	}
	defer jsonFile.Close()

	_, err = jsonFile.Write(jsonData)
	if err != nil {
		fmt.Println("Error writing JSON to file:", err)
		return
	}

	fmt.Println("JSON data has been written to output.json")
}
