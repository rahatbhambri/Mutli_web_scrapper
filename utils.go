package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"

	goquery "github.com/PuerkitoBio/goquery"
)

func AppendToFile(filepath string, data string) error {
	// Open the file for appending, create if it doesn't exist
	file, err := os.OpenFile(filepath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return err
	}
	defer file.Close()
	// Create a buffered writer from the file
	writer := bufio.NewWriter(file)
	_, err = writer.WriteString(data)
	if err != nil {
		return err
	}
	// Ensure all buffered data is written to the file
	err = writer.Flush()
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
	// return if webpage is unavailable
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
