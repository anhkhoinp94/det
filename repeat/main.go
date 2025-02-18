package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"strings"
	"time"
)

const (
	white  = "\033[38;5;237m"
	yellow = "\033[33m"
	reset  = "\033[0m"
)

type Paragraph struct {
	P string `json:"p"`
}

func main() {
	file, err := os.Open("paragraphs.json")
	if err != nil {
		fmt.Println("Error opening JSON file:", err)
		return
	}
	defer file.Close()

	var paragraphs []Paragraph
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&paragraphs); err != nil {
		fmt.Println("Error decoding JSON file:", err)
		return
	}

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(paragraphs), func(i, j int) {
		paragraphs[i], paragraphs[j] = paragraphs[j], paragraphs[i]
	})

	clearScreen()

	for i := 0; i < len(paragraphs); i++ {
		text := paragraphs[i].P

		sentences := strings.Split(text, "\n\n")
		if len(sentences) > 0 && sentences[len(sentences)-1] == "" {
			sentences = sentences[:len(sentences)-1] // Remove empty last element
		}

		// Speak each sentence 10 times
		for _, sentence := range sentences {
			sentence = strings.TrimSpace(sentence)
			// fmt.Println(sentence)
			printText(sentence, sentences)
			for j := 0; j < 2; j++ {
				repeatedTime := speak(sentence)
				time.Sleep(repeatedTime)
			}
			clearScreen()
		}
	}

	fmt.Println("All paragraphs have been repeated. Exiting.")
}

func printText(sentence string, sentences []string) {
	for s := 0; s < len(sentences); s++ {
		if strings.Contains(sentences[s], sentence) {
			fmt.Printf("%s%s%s\n", yellow, sentences[s], reset)
		} else {
			fmt.Printf("%s%s%s\n", white, sentences[s], reset)
		}
	}
}

func clearScreen() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func speak(text string) time.Duration {
	start := time.Now()
	cmd := exec.Command("tts", text)
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error putting the system to sleep:", err)
		return 0
	}
	elapsed := time.Since(start)
	return elapsed
}
