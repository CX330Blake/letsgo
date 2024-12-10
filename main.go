package main

import (
	"bufio"
	"flag"
	"io"
	"log"
	"os"

	"github.com/CX330Blake/letsgo/pkg/greet"
	"github.com/CX330Blake/letsgo/pkg/letsgo"
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

func main() {

	log.SetOutput(io.Discard)
	// Parse command line arguments
	url := flag.String("url", "", "Target URL (e.g. http://example.com)")
	param := flag.String("param", "file", "Parameter for testing (default is 'file')")
	wordlistFile := flag.String("wordlist", "default.txt", "Wordlist path")
	fileRoot := flag.String("root", "", "Root of the server file (e.g. https://example.com/image?filename=/var/www/images/1337.jpg, then root is `/var/www/images`, don't need to include the last `/`)")
	extension := flag.String("extension", "", "File extension (e.g. jpg, png, txt, etc.), this will triger the null byte bypass mode")

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

	letsgo.Test(*url, *param, *fileRoot, *extension, payloads)
	greet.End()
}
