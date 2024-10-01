package router

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

// Test when the file exists and is successfully read
func TestGetDataHandler_Success(t *testing.T) {
	// Create a temporary file with content
	testFilePath := "static/web_data.txt"
	err := os.MkdirAll("static", 0755) // Ensure the directory exists
	if err != nil {
		t.Fatalf("Error creating static directory: %v", err)
	}

	err = ioutil.WriteFile(testFilePath, []byte("Hello, World!"), 0644)
	if err != nil {
		t.Fatalf("Error writing test file: %v", err)
	}

	// Create a new HTTP request
	req, err := http.NewRequest("GET", "/getData", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a response recorder (to capture the response)
	rr := httptest.NewRecorder()

	// Call the handler
	handler := http.HandlerFunc(GetDataHandler)
	handler.ServeHTTP(rr, req)

	// Check if the status code is 200 OK
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v, want %v", status, http.StatusOK)
	}

	// Check if the response body is the expected file content
	expected := "Hello, World!"
	if rr.Body.String() != expected {
		t.Errorf("Handler returned unexpected body: got %v, want %v", rr.Body.String(), expected)
	}

	// Clean up the test file
	os.Remove(testFilePath)
}

// Test when the file does not exist
func TestGetDataHandler_FileNotFound(t *testing.T) {
	// Remove the file if it exists
	testFilePath := "static/web_data.txt"
	os.Remove(testFilePath)

	// Create a new HTTP request
	req, err := http.NewRequest("GET", "/getData", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a response recorder
	rr := httptest.NewRecorder()

	// Call the handler
	handler := http.HandlerFunc(GetDataHandler)
	handler.ServeHTTP(rr, req)

	// Check if the status code is 500 Internal Server Error
	if status := rr.Code; status != http.StatusInternalServerError {
		t.Errorf("Handler returned wrong status code: got %v, want %v", status, http.StatusInternalServerError)
	}

	// Check if the response body contains the error message
	expected := "Unable to read data file\n"
	if rr.Body.String() != expected {
		t.Errorf("Handler returned unexpected body: got %v, want %v", rr.Body.String(), expected)
	}
}

// Test if content-type is set correctly
func TestGetDataHandler_ContentType(t *testing.T) {
	// Create a temporary file with content
	testFilePath := "static/web_data.txt"
	err := ioutil.WriteFile(testFilePath, []byte("Sample content"), 0644)
	if err != nil {
		t.Fatalf("Error writing test file: %v", err)
	}

	// Create a new HTTP request
	req, err := http.NewRequest("GET", "/getData", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a response recorder
	rr := httptest.NewRecorder()

	// Call the handler
	handler := http.HandlerFunc(GetDataHandler)
	handler.ServeHTTP(rr, req)

	// Check if the Content-Type is text/plain
	if contentType := rr.Header().Get("Content-Type"); contentType != "text/plain" {
		t.Errorf("Content-Type header is wrong: got %v, want %v", contentType, "text/plain")
	}

	// Clean up the test file
	os.Remove(testFilePath)
}
