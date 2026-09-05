package main

import (
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strings"
)

// Accepts the quit channel and returns the output channels
func generateUrls(quit <-chan int) <-chan string { 
	// create the output channel
	urls := make(chan string)
	go func ()  {
		defer close(urls) // once complete closes the output channels
		for i := 100; i <= 130; i++ {
			url := fmt.Sprintf("https://rfc-editor.org/rfc/rfc%d.txt", i)
			select {
			case urls <- url: // writes 50 urls to the output channels
			case <-quit:
				return 
			}
		}
	}()
	return urls // return the output channel
}

func downloadPages(quit <-chan int, urls <-chan string) <-chan string {
	pages := make(chan string) // Creates the output channel, shich will contain downloaded web pages
	go func() {
		// Close the output channel when its finished
		defer close(pages)
		moreData, url := true, ""
		// continues selecting if theres more data on the input channel
		for moreData {
			select {
			// updates variables with a new message and flag
			//  to show whether there is more data
			case url, moreData = <-urls:
				if moreData {
					resp, _ := http.Get(url)
					if resp.StatusCode != 200 {
						panic("Server error: " + resp.Status)
					}
					body, _ := io.ReadAll(resp.Body)
					pages <- string(body)
					resp.Body.Close()
				}
			case <-quit: //when a message arrives on the quit channel, terminate the goroutine
				return 
			}
		}
	}()
	return pages // return the output channel
}

func extactWords(quit <-chan int, pages <-chan string) <-chan string {
	// Creates the output channel, which will contain extracted words
	words := make(chan string)
	go func ()  {
		defer close(words)	
		wordRegex := regexp.MustCompile(`[a-zA-Z]+`) // creates a regular expresion to extract the words
		moreData, pg := true, ""
		for moreData {
			select {
			case pg, moreData = <-pages: // updates variables with a new message and flag to show whether there is more data
				if moreData {
					for _, word := range wordRegex.FindAllString(pg, -1) {
						words <- strings.ToLower(word)
					}
				}
			case <-quit: // When a message arrives on the quit channel, terminates goroutine
			return 	
			}
		}
	}()
	return words // return the output channel
}

func main() {
	quit := make(chan int)
	defer close(quit)
	// adds the new goroutine that downloads pages to the existing pipeline
	results := extactWords(quit, downloadPages(quit, generateUrls(quit)))
	for results := range results {
		fmt.Println(results)
	}
}