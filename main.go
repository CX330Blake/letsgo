package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"sync"

	"github.com/CX330Blake/letsgo/pkg/greet"
	"github.com/fatih/color"
)

// 加載 Payload
func loadPayloads(filePath string) ([]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("Cannot load the wordlist: %v", err)
	}
	defer file.Close()

	var payloads []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		payloads = append(payloads, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("Error on reading the wordlist: %v", err)
	}

	return payloads, nil
}

// Send request and check response
func testPayload(url string, param string, payload string) {
	fullURL := fmt.Sprintf("%s?%s=%s", url, param, payload)
	resp, err := http.Get(fullURL)
	if err != nil {

		color.Red("[!] Request failed: %v\n", err)
		return
	}
	defer resp.Body.Close()

	// Simple check if the response might indicate a vulnerability
	if resp.StatusCode == http.StatusOK {
		color.Green("[+] Found Possible Vuln URL: %s (Status: %d)\n", fullURL, resp.StatusCode)
	}
}

// Multi-threaded testing
func testPathTraversal(url string, param string, payloads []string) {
	_, err := http.Get(url)
	if err != nil {
		color.Red("[!] Host seems down...\n")
		return
	}

	var wg sync.WaitGroup

	for _, payload := range payloads {
		wg.Add(1)
		go func(pl string) {
			defer wg.Done()
			testPayload(url, param, pl)
		}(payload)
	}

	wg.Wait()
}

func main() {

	log.SetOutput(io.Discard)
	// Parse command line arguments
	url := flag.String("url", "", "Target URL (E.g. http://example.com)")
	param := flag.String("param", "file", "Parameter for testing (default is 'file')")
	payloadFile := flag.String("payloads", "payloads.txt", "Wordlist path")
	flag.Parse()

	if *url == "" {
		color.Yellow("Usage: ./LetsGo -url <https://example.com>")
		os.Exit(1)
	}

	// Load payload
	payloads, err := loadPayloads(*payloadFile)
	if err != nil {
		color.Red("[!] Cannot laod the wordlist: %v\n", err)
		os.Exit(1)
	}

	// Hack the planet
	greet.Hello()
	testPathTraversal(*url, *param, payloads)
	greet.End()
}
