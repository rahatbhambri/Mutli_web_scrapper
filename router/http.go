package router

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// Handler for the /getData route
func GetDataHandler(w http.ResponseWriter, r *http.Request) {
	// Define the path to the static file
	filePath := "static/web_data.txt"

	// Read the content from the file
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		// Handle error if the file can't be read
		http.Error(w, "Unable to read data file", http.StatusInternalServerError)
		fmt.Println("Error reading file:", err)
		return
	}

	// Send the file content as the response
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}
