package main

import (
	"fmt"
	"math/rand"
	"os/exec"
	"sort"

	"sleep/det"
	"time"
)

func main() {
	// sleepPc(35)
	// shuffle()
	det.Convert()
}

func shuffle() {
	slice := []int{21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	// slice := []int{22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38}
	fmt.Println("Trước khi trộn:", slice)
	rand.Seed(time.Now().UnixNano()) // Khởi tạo seed ngẫu nhiên
	for i := len(slice) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)                   // Lấy chỉ số ngẫu nhiên trong khoảng [0, i]
		slice[i], slice[j] = slice[j], slice[i] // Hoán đổi 2 phần tử
	}
	fmt.Println("Sau khi trộn:", slice)
	// Sort the first 20 elements
	first20 := slice[:20]
	sort.Ints(first20)
	last20 := slice[20:]
	sort.Ints(last20)
	fmt.Println("Sau khi trộn phần đầu:", first20)
	fmt.Println("Phần còn lại:", last20)
}

func sleepPc(minutes int) {
	for i := 1; i <= minutes; i++ {
		fmt.Printf("Go to sleep in the next %d minutes\n", minutes+1-i)
		time.Sleep(1 * time.Minute)
	}

	// Execute the command to put the computer to sleep
	cmd := exec.Command("rundll32.exe", "powrprof.dll,SetSuspendState", "0", "1", "0")
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error putting the system to sleep:", err)
		return
	}

	fmt.Println("The system is now sleeping.")
}
