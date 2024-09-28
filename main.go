package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"
	"webscrapper/router"
	"webscrapper/utils"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func updateFile() {
	var wg sync.WaitGroup
	filepath := "static/web_data.txt"
	_, err := os.Create(filepath)
	if err != nil {
		fmt.Println("Error creating or replacing file:", err)
		return
	}
	ch := make(chan string, 4)
	urls := []string{"https://en.wikipedia.org/wiki/Tiger", "http://www.facebook.com", "http://www.cnet.com", "https://en.wikipedia.org/wiki/Taj_Mahal"}
	// spawn gouroutines to start scrapping
	for _, url := range urls {
		go utils.StartScrapping(url, ch, &wg)
		wg.Add(1)
	}
	wg.Wait()
	close(ch)

	for data := range ch {
		err := utils.AppendToFile(filepath, data)
		if err != nil {
			fmt.Println("error writing to file")
		} else {
			fmt.Println("Successfully written to file")
		}
	}
	fmt.Println("All done")
}

func main() {
	r := chi.NewRouter()

	// Use some middleware (like request logging)
	r.Use(middleware.Logger)

	// Define the /getData route
	r.Get("/getData", router.GetDataHandler)

	// Start the server on port 8080
	fmt.Println("Server starting at http://localhost:8080...")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal("Server failed to start:", err)
	}
}
