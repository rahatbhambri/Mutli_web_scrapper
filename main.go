package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	filepath := "web_data.txt"
	_, err := os.Create(filepath)
	if err != nil {
		fmt.Println("Error creating or replacing file:", err)
		return
	}
	startTime := time.Now()
	ch := make(chan string, 3)
	// spawn gouroutines to start scrapping
	go StartScrapping("https://en.wikipedia.org/wiki/Tiger", ch)
	go StartScrapping("http://www.facebook.com", ch)
	go StartScrapping("http://www.cnet.com", ch)
	go StartScrapping("https://en.wikipedia.org/wiki/Taj_Mahal", ch)
	var local_data []string
	for {
		// Check elapsed time since start
		elapsed := time.Since(startTime)
		// Exit loop if 10 seconds have passed
		if elapsed >= 10*time.Second {
			break
		}
		// Example: Simulate receiving values from a channel (replace with actual logic)
		select {
		case val := <-ch:
			if val != "" {
				local_data = append(local_data, val)
			}
		default:
			// No value received, continue loop or perform other operations
		}
	}
	for _, data := range local_data {
		err := AppendToFile(filepath, data)
		if err != nil {
			fmt.Println("error writing to file")
		} else {
			fmt.Println("Successfully written to file")
		}
	}
}
