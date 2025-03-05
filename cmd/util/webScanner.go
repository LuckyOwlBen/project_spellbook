package util

import (
	"fmt"
	"net/http"
	"os"
	"strings"
)

var smallWebDirectoryPath = "/home/demeter/Downloads/SecLists/Discovery/Web-Content/directory-list-2.3-small.txt"

type ScanResult struct {
	StatusCode int
	URL        string
}

func ScanWebDirectories(ipAddress string, port string) {
	var results []ScanResult
	wordList := parseWordList(smallWebDirectoryPath)
	for _, word := range wordList {
		if word == "" {
			continue
		}
		url := fmt.Sprintf("http://%s:%s/%s", ipAddress, port, word)
		fmt.Println("Scanning", url)
		resp, err := http.Get(url)
		if err != nil {
			//fmt.Println("Error scanning", url, err)
			continue
		}
		results = append(results, ScanResult{
			StatusCode: resp.StatusCode,
			URL:        url,
		})
	}
	for _, result := range results {
		fmt.Println("URL:", result.URL, "Status Code:", result.StatusCode)
	}
}

func parseWordList(secListPath string) []string {
	unparsedWordList, err := os.ReadFile(secListPath)
	if err != nil {
		panic(err)
	}
	return strings.Split(string(unparsedWordList), "\n")
}
