package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

const (
	white  = "\033[38;5;237m"
	black  = "\033[38;5;232m"
	yellow = "\033[33m"
	reset  = "\033[0m"
)

type Paragraph struct {
	P string `json:"p"`
}

func main() {
	file, err := os.Open("stories.json")
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

	// rand.Seed(time.Now().UnixNano())
	// rand.Shuffle(len(paragraphs), func(i, j int) {
	// 	paragraphs[i], paragraphs[j] = paragraphs[j], paragraphs[i]
	// })

	// paragraphs = reverseSlice(paragraphs)

	for i := 0; i < len(paragraphs); i++ {
		clearScreen()
		text := paragraphs[i].P

		sentences := strings.Split(text, "\n\n")
		if len(sentences) > 0 && sentences[len(sentences)-1] == "" {
			sentences = sentences[:len(sentences)-1]
		}

		for _, sentence := range sentences {
			sentence = strings.TrimSpace(sentence)
			// fmt.Println(sentence)
			printText(sentence, sentences)
			speak(sentence)
			clearScreen()
		}

		fmt.Printf("%v paragraphs left\n", len(paragraphs)-i-1)
		// waitForEnter()
	}

	fmt.Println("All paragraphs have been listened. Exiting.")
}

//lint:ignore U1000 This function is used to reverse slice
func reverseSlice(slice []Paragraph) []Paragraph {
	for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}
	return slice
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

//lint:ignore U1000 This function is used to wait for enter
func waitForEnter() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Press Enter to continue...")
	reader.ReadString('\n')
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
