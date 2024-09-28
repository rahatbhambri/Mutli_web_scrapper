package utils

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"sync"
	"testing"
)

func TestAppendToFile(t *testing.T) {
	// Test setup: create a temporary file
	tmpfile, err := ioutil.TempFile("", "testfile")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(tmpfile.Name()) // Clean up

	// Write initial content to the file
	initialContent := "Initial content\n"
	err = AppendToFile(tmpfile.Name(), initialContent)
	if err != nil {
		t.Fatalf("Failed to append initial content: %v", err)
	}

	// Append new content
	newContent := "New content\n"
	err = AppendToFile(tmpfile.Name(), newContent)
	if err != nil {
		t.Fatalf("Failed to append new content: %v", err)
	}

	// Read the file to verify the content
	content, err := ioutil.ReadFile(tmpfile.Name())
	if err != nil {
		t.Fatalf("Failed to read file: %v", err)
	}

	expectedContent := initialContent + newContent
	if string(content) != expectedContent {
		t.Errorf("File content mismatch: expected %q, got %q", expectedContent, string(content))
	}
}

func TestAppendToFileError(t *testing.T) {
	// Test case where the file cannot be opened (e.g., directory does not exist)
	err := AppendToFile("/invalid/path/testfile", "Some content")
	if err == nil {
		t.Errorf("Expected error when opening a file in an invalid path, but got nil")
	}
}

func TestStartScrapping(t *testing.T) {
	// Create a mock server to simulate the webpage
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "text/html")
		w.Write([]byte(`
			<html>
				<body>
					<p>This is paragraph 1.</p>
					<p>This is paragraph 2.</p>
				</body>
			</html>
		`))
	}))
	defer mockServer.Close()

	// Set up the channel and WaitGroup
	ch := make(chan string, 1)
	var wg sync.WaitGroup
	wg.Add(1)

	// Call the StartScrapping function
	go StartScrapping(mockServer.URL, ch, &wg)
	wg.Wait()

	// Verify the result from the channel
	result := <-ch
	expected := "This is paragraph 1.This is paragraph 2."
	if result != expected {
		t.Errorf("Scrapped data mismatch: expected %q, got %q", expected, result)
	}
}

func TestStartScrappingError(t *testing.T) {
	// Simulate a server returning a 404 status code
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Not Found", http.StatusNotFound)
	}))
	defer mockServer.Close()

	// Set up the channel and WaitGroup
	ch := make(chan string, 1)
	var wg sync.WaitGroup
	wg.Add(1)

	// Call the StartScrapping function with a server returning 404
	go StartScrapping(mockServer.URL, ch, &wg)
	wg.Wait()

	// Ensure that nothing was sent to the channel
	select {
	case result := <-ch:
		t.Errorf("Expected no result, but got %q", result)
	default:
		// Success: no data was sent to the channel
	}
}
