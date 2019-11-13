package main

import (
      "os"
      "github.com/fatih/color"
      "net/http"
      "flag"
      "bufio"
)

var scanURL = flag.String("url", "https://google.com", "URL of the site you wish to check.")
var headerList = flag.String("file", "headers.txt", "Custom headers file to read from.")

func main() {

    // Parse Config Flags
    flag.Parse()

    // Get URL Response + Headers
    resp, err := http.Get(*scanURL)
    if err != nil {
        color.Red("Failed to GET "+*scanURL+" please check to ensure this URL exists.")
        panic(err)
    }
    defer resp.Body.Close()

    // Read Security Headers from File
    headers, err := scanLines(*headerList)
    if err != nil {
        color.Red("Unable to open "+*headerList+" please check to ensure this file exists.")
        panic(err)
    }

    // Find Security Headers That Match
    // headers = slice
    // resp.Header = map
    // fmt.Println(headers)
    // fmt.Println("")
    // fmt.Println(resp.Header)

    for _, header := range headers {
        _, ok := resp.Header[header]
        if ok {
          color.HiGreen(header + " - Found")
        } else {
          color.Red(header + " - Not Found")
        }
    }
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
