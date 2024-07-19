package main

import (
	"fmt"
	"net/http"

	goquery "github.com/PuerkitoBio/goquery"
)

func main() {
	fmt.Println("Hello")

	ch := make(chan string, 3)
	go StartScrapping("http://www.wikipedia.com", ch)
	go StartScrapping("http://www.facebook.com", ch)
	go StartScrapping("http://www.instagram.com", ch)
	go StartScrapping("http://www.google.com", ch)

	for {
		val := <-ch
		fmt.Println(val)
	}

}

func StartScrapping(url string, ch chan string) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	// fmt.Println(resp)

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
		doc.Find("p").Each(func(i int, s *goquery.Selection) {
			S += s.Text()
		})
	}
	// fmt.Println("at end")
	ch <- S

}
