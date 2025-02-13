package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
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

	// rand.Seed(time.Now().UnixNano())
	// rand.Shuffle(len(paragraphs), func(i, j int) {
	// 	paragraphs[i], paragraphs[j] = paragraphs[j], paragraphs[i]
	// })
	paragraphs = reverseSlice(paragraphs)

	for i := 0; i < len(paragraphs); i++ {
		clearScreen()
		text := paragraphs[i].P
		fmt.Println(text)
		fmt.Printf("%v paragraphs left\n", len(paragraphs)-i-1)
		time.Sleep(1 * time.Second)
		speak(text)
		time.Sleep(1 * time.Second)
		speak(text)
		waitForEnter()
	}

	fmt.Println("All paragraphs have been listened. Exiting.")
}

func reverseSlice(slice []Paragraph) []Paragraph {
	for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}
	return slice
}

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
