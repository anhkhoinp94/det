package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"strconv"
	"time"
)

type Question struct {
	Question string `json:"question"`
}

func loadQuestions(file string) ([]Question, error) {
	var questions []Question
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(data, &questions); err != nil {
		return nil, err
	}
	return questions, nil
}

func main() {
	questions, err := loadQuestions("questions.json")
	if err != nil {
		fmt.Println("Error loading questions:", err)
		return
	}

	rand.Seed(time.Now().UnixNano())
	usedIndexes := make(map[int]bool)

	colors := []string{
		"\033[31m", // Red
		"\033[32m", // Green
		"\033[33m", // Yellow
		"\033[34m", // Blue
		"\033[35m", // Magenta
		"\033[36m", // Cyan
	}

	reset := "\033[0m"

	for {
		if len(usedIndexes) == len(questions) {
			fmt.Println("No more questions left.")
			break
		}

		var index int
		for {
			index = rand.Intn(len(questions))
			if !usedIndexes[index] {
				usedIndexes[index] = true
				break
			}
		}

		color := colors[rand.Intn(len(colors))]
		fmt.Println(color + `(Remaining: ` + strconv.Itoa(len(questions)-len(usedIndexes)) + `)` + reset)
		fmt.Println(color + questions[index].Question + reset)

		var input string
		fmt.Print("Press 'y' or 'q': ")
		fmt.Scanln(&input)
		if input == "q" {
			break
		}
	}
}
