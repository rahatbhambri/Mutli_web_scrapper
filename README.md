# Multi Web Scraper

## Overview
Multi Web Scraper is a concurrent web scraping application built using Go. It allows users to scrape multiple URLs simultaneously for text content and store the extracted data in a file. The stored data can later be used to train AI models and build a question-answering system on top of it.

## Features
- **Concurrent Scraping:** Scrapes multiple URLs simultaneously for efficiency.
- **Data Storage:** Extracted text is stored in a file (`web_data.txt`).
- **AI Training Ready:** The stored text can be used for AI-based applications like question-answering systems.
- **Lightweight & Fast:** Built using Go for high performance and minimal resource consumption.

## Tech Stack
- **Programming Language:** Go (Golang)
- **Libraries Used:** net/http, goquery (for HTML parsing)

## Installation
```sh
git clone https://github.com/rahatbhambri/Mutli_web_scrapper.git
cd Mutli_web_scrapper
```

## Usage
1. Run the scraper using the following command:
   ```sh
   go run main.go
   ```
2. The extracted text data will be stored in `web_data.txt`.
![image](https://github.com/user-attachments/assets/0dab627d-5ff9-4a01-b25f-9c3f0a61db31)

## Future Enhancements
- Implement a database for structured storage.
- Support for different data formats (JSON, CSV).
- Build a question-answering system on top of the collected data.
- Add support for scraping images and metadata.

## Contributing
Contributions are welcome! Feel free to submit issues and pull requests.

## License
(Include license details if applicable.)





