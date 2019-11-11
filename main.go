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
   
    // for _, header := range headers {
    //     fmt.Println(header)
    // }

    // Find Security Headers That Match
    for _, header := range headers {
        var found = ""
        for k, _ := range resp.Header {
            // fmt.Println(header)
            if header == k {
                found := "OK"
            } else {
                found := "MISSING"
            }
        }
        color.HiGreen(header + " %d\n",found)
        // fmt.Print(k)
        // fmt.Print(" : ")
        // fmt.Println(v)
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
