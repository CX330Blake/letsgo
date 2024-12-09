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

// Load payload
func loadPayloads(filePath string) ([]string, bool) {
	useDefault := false
	file, err := os.Open(filePath)
	if err != nil {
		file, err = os.Open("./default.txt")
		useDefault = true
		// return nil, fmt.Errorf("cannot load the wordlist: %v", err)
	}
	defer file.Close()

	var payloads []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		payloads = append(payloads, scanner.Text())
	}

	// if err := scanner.Err(); err != nil {
	// 	return nil, useDefault
	// }

	return payloads, useDefault
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
	wordlistFile := flag.String("wordlist", "default.txt", "Wordlist path")
	flag.Parse()

	if *url == "" {
		color.Yellow("Usage: ./LetsGo -url <https://example.com>")
		os.Exit(1)
	}

	// Hack the planet
	greet.Hello()

	payloads, useDefault := loadPayloads(*wordlistFile)
	if useDefault && *wordlistFile != "default.txt" {
		color.Magenta("[!] Cannot load the wordlist, using default list now...\n")
	} else if *wordlistFile == "default.txt" {
		color.Magenta("[+] Using default wordlist...\n")
	}

	testPathTraversal(*url, *param, payloads)
	greet.End()
}
