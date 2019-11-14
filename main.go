package main

import (
	"bufio"
	"flag"
	"github.com/fatih/color"
	"net/http"
	"os"
	"strings"
)

var scanURL = flag.String("url", "google.com", "URL of the site you wish to check.")
var scanScope = flag.String("scope", "", "Custom scope file to read from.")
var headerList = flag.String("file", "headers.txt", "Custom headers file to read from.")

func main() {

	// Parse Config Flags
	flag.Parse()

	// Read URL Scope from File
	if isFlagPassed(*scanScope); *scanScope != "" {
		color.Red("SCOPE")
		scopes, err := scanLines(*scanScope)
		if err != nil {
			color.Red("Unable to open " + *scanScope + " please check to ensure this file exists.")
			panic(err)
		}
		for _, scope := range scopes {
			scanSite(scope)
			continue
		}
	}

	// Find Security Headers That Match
	// headers = slice
	// resp.Header = map
	// fmt.Println(headers)
	// fmt.Println("")
	// fmt.Println(resp.Header)
	if isFlagPassed(*scanURL); *scanURL != "google.com" {
		color.Red("URL")
		scanSite(*scanURL)
	}
}

func scanSite(url string) ([]string, error) {
	// Read Security Headers from File
	headers, err := scanLines(*headerList)
	if err != nil {
		color.Red("Unable to open " + *headerList + " please check to ensure this file exists.")
		panic(err)
	}

	// Get URL Response + Headers
	resp, err := http.Get("https://" + url)
	if err != nil {
		color.Red("Failed to GET " + *scanURL + " please check to ensure this URL exists.")
		panic(err)
	}
	defer resp.Body.Close()

	color.Yellow("Domain: https://" + url)
	for _, header := range headers {
		_, ok := resp.Header[header]
		if ok {
			// Convert []string to string
			value := strings.Join(resp.Header[header], " ")
			// Print the found header
			color.Green("[✔] " + header + " (" + value + ")")
			// Instead of using else we cano continue the loop before moving out of the loop
			continue
		}
		color.Red("[✗] " + header)
	}
	return nil, nil
}

func scanLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, nil
}

func isFlagPassed(name string) bool {
	found := false
	flag.Visit(func(f *flag.Flag) {
		if f.Name == name {
			found = true
		}
	})
	return found
}
