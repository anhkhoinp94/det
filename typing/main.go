package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"strings"
	"time"
	"unicode/utf8"

	"golang.org/x/crypto/ssh/terminal"
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
	usedIndices := make(map[int]bool)

	for len(usedIndices) < len(paragraphs) {
		var index int
		for {
			index = rand.Intn(len(paragraphs))
			if !usedIndices[index] {
				usedIndices[index] = true
				break
			}
		}

		text := paragraphs[index].P
		input := ""

		clearScreen()
		printText(text, input)

		go func() {
			speak(text)
		}()
		go func() {
			time.Sleep(180 * time.Second)
			speak(text)
		}()

		// Set the terminal to raw mode to capture each keystroke
		oldState, err := terminal.MakeRaw(int(os.Stdin.Fd()))
		if err != nil {
			fmt.Println("Error setting raw mode:", err)
			return
		}
		defer terminal.Restore(int(os.Stdin.Fd()), oldState)

		reader := bufio.NewReader(os.Stdin)

		for {
			char, _, err := reader.ReadRune()
			if err != nil {
				fmt.Println("Error reading input:", err)
				return
			}

			if utf8.RuneCountInString(input) < utf8.RuneCountInString(text) && rune(text[utf8.RuneCountInString(input)]) == char {
				input += string(char)
			} else {
				fc := string(rune(text[utf8.RuneCountInString(input)]))
				if fc == "." || fc == "," || fc == "'" {
				} else {
					items := strings.Split(input, " ")
					if len(items) > 1 {
						input = strings.Join(items[:len(items)-1], " ") + " "
					} else {
						input = ""
					}
				}
			}

			clearScreen()
			printText(text, input)

			if input == text {
				fmt.Println("Congratulations! You've typed the text correctly.")
				fmt.Printf("%v paragraphs left\n", len(paragraphs)-len(usedIndices))
				break
			}
		}

		fmt.Println("Press any key to continue to the next paragraph...")
		_, _, _ = reader.ReadRune() // Wait for any key press
	}

	fmt.Println("All paragraphs have been typed. Exiting.")
}

func printText(text, input string) {
	boldText := text[:utf8.RuneCountInString(input)]
	lightText := text[utf8.RuneCountInString(input):]

	fmt.Printf("%s%s%s%s%s\n", yellow, boldText, white, lightText, reset)
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
