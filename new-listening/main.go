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
			sentence = strings.TrimSpace(sentence) // Ensure it ends with a period
			fmt.Println(sentence)
			for j := 0; j < 5; j++ {
				speak(sentence)
				time.Sleep(50 * time.Millisecond)
			}
			clearScreen()
		}

		// Speak the whole paragraph 5 times
		fmt.Println(text)
		fmt.Printf("%v paragraphs left\n", len(paragraphs)-i-1)
		for k := 0; k < 10; k++ {
			speak(text)
			time.Sleep(50 * time.Millisecond)
		}
		clearScreen()
	}

	fmt.Println("All paragraphs have been listened. Exiting.")
}

func clearScreen() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func speak(text string) {
	cmd := exec.Command("tts", text)
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error putting the system to sleep:", err)
		return
	}
}
