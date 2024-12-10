package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"sync"

	"github.com/CX330Blake/letsgo/pkg/greet"
	"github.com/fatih/color"
)

// Load payload
func loadPayloads(filePath string) ([]string, bool) {
	useDefault := false
	file, err := os.Open(filePath)
	if err != nil {
		file, _ = os.Open("./default.txt")
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
func testPayload(url string, param string, fileRoot string, wordlist string) {
	if !strings.HasSuffix(fileRoot, "/") {
		fileRoot += "/"
	}

	fullURL := fmt.Sprintf("%s?%s=%s%s", url, param, fileRoot, wordlist)
	resp, err := http.Get(fullURL)
	if err != nil {
		// color.Red("[!] Request failed: %v\n", err)
		return
	}
	defer resp.Body.Close()

	// Simple check if the response might indicate a vulnerability
	if resp.StatusCode == http.StatusOK {
		color.Green("[+] Found Possible Vuln URL: %s (Status: %d)\n", fullURL, resp.StatusCode)
	}
}

// Multi-threaded testing
func testPathTraversal(url string, param string, fileRoot string, wordlist []string) {
	_, err := http.Get(url)
	if err != nil {
		color.Red("[!] Host seems down...\n")
		return
	}

	var wg sync.WaitGroup

	for _, payload := range wordlist {
		wg.Add(1)
		go func(wl string) {
			defer wg.Done()
			testPayload(url, param, fileRoot, wl)
		}(payload)
	}

	wg.Wait()
}

func main() {

	log.SetOutput(io.Discard)
	// Parse command line arguments
	url := flag.String("url", "", "Target URL (e.g. http://example.com)")
	param := flag.String("param", "file", "Parameter for testing (default is 'file')")
	wordlistFile := flag.String("wordlist", "default.txt", "Wordlist path")
	fileRoot := flag.String("root", "", "Root of the server file (e.g. https://example.com/image?filename=/var/www/images/1337.jpg, then root is `/var/www/images`, don't need to include the last `/`)")

	flag.Parse()

	if *url == "" {
		color.Magenta("Basic usage: ./letsgo -url <https://example.com>")
		os.Exit(1)
	}

	// Hack the planet
	greet.Hello()

	payloads, useDefault := loadPayloads(*wordlistFile)
	if useDefault && *wordlistFile != "default.txt" {
		color.Magenta("[*] Cannot load the wordlist, using default list now...\n")
	} else if *wordlistFile == "default.txt" {
		color.Magenta("[+] Using default wordlist...\n")
	}

	testPathTraversal(*url, *param, *fileRoot, payloads)
	greet.End()
}
