package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {
	duration := 2 * time.Minute // Set the countdown duration here
	resetDuration := duration
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	endTime := time.Now().Add(duration)
	colors := []string{
		"\033[32m", // Green
		"\033[33m", // Yellow
		"\033[31m", // Red
		"\033[34m", // Blue
		"\033[35m", // Magenta
		"\033[36m", // Cyan
	}
	reset := "\033[0m"
	colorIndex := 0

	// Goroutine to listen for Enter key press
	go func() {
		reader := bufio.NewReader(os.Stdin)
		for {
			_, _ = reader.ReadString('\n')
			endTime = time.Now().Add(resetDuration)
		}
	}()

	for {
		select {
		case <-ticker.C:
			remaining := endTime.Sub(time.Now())
			if remaining <= 0 {
				fmt.Println("Time's up!")
				return
			}

			if int(remaining.Seconds())%60 == 0 {
				colorIndex = (colorIndex + 1) % len(colors)
			}

			fmt.Print("\033[H\033[2J") // Clear the terminal
			fmt.Printf("Remaining time: %s%s%s\n", colors[colorIndex], fmtDuration(remaining), reset)
		}
	}
}

func fmtDuration(d time.Duration) string {
	h := int(d.Hours())
	m := int(d.Minutes()) % 60
	s := int(d.Seconds()) % 60
	return fmt.Sprintf("%02d:%02d:%02d", h, m, s)
}
