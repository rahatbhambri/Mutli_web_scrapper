package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"time"

	goquery "github.com/PuerkitoBio/goquery"
)

func main() {
	// fmt.Println("Hello")
	filepath := "web_data.txt"
	startTime := time.Now()

	ch := make(chan string, 3)
	go StartScrapping("https://en.wikipedia.org/wiki/Tiger", ch)
	go StartScrapping("http://www.facebook.com", ch)
	go StartScrapping("http://www.instagram.com", ch)
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
		err := WriteToFile(filepath, data)
		if err != nil {
			fmt.Println("error writing to file")
		} else {
			fmt.Println("Successfully written to file")
		}
	}
}

func WriteToFile(filepath string, data string) error {
	// Open the file for writing, create if it doesn't exist, truncate if it does
	file, err := os.Create(filepath)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return err
	}
	defer file.Close()

	// Create a buffered writer from the file
	writer := bufio.NewWriter(file)
	_, err = writer.WriteString(data)
	if err != nil {
		return err
	}
	return nil
}

func StartScrapping(url string, ch chan string) {
	// Getting data from webpage
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return
	}

	S := ""
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		fmt.Println("error fetching from website")
		return
	} else {
		doc, err := goquery.NewDocumentFromReader(resp.Body)
		if err != nil {
			fmt.Println(err)
			return
		}
		// Getting all text data from webpages
		doc.Find("p").Each(func(i int, s *goquery.Selection) {
			S += s.Text()
		})
	}
	ch <- S
}
