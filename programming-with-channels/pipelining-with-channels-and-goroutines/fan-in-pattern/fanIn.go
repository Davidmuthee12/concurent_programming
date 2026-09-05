package main

import (
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strings"
	"sync"
)

const downloaders = 20

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

func extractWords(quit <-chan int, pages <-chan string) <-chan string {
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

func FanIn[k any](quit <-chan int, allChannels ...<-chan k) chan k {
	// creates a waitgroup, setting the size to be equalto the number of input channels
	wg := sync.WaitGroup{}
	wg.Add(len(allChannels))
	// Creates the input channel
	output := make(chan k)
	for _, c := range allChannels {
		go func (channel <-chan k)  { // Starts a goroutine for every input channel
			// once the goroutine terminates, marks the waitgroup as done
			defer wg.Done()
			for i := range channel {
				select {
				case output <- i: // Forwards each received message to the shared output channel
				case <-quit: // if quit channel is closed, terminates the goroutine
					return 
				}
			}
		}(c) // passes one input channel to the goroutine
	}
	// waits for all goroutines to finish and then closes the output channel
	go func ()  {
		wg.Wait()
		close(output)
	}()
	return output
}

func main() {
	quit := make(chan int)
	defer close(quit)
	urls := generateUrls(quit)
	pages := make([]<-chan string, downloaders)
	for i := 0; i < downloaders; i++ {
		pages[i] = downloadPages(quit, urls)
	}
	results := extractWords(quit, FanIn(quit, pages...)) // joins all the pages channels into one channel using the fan-in pattern
	for result := range results {
		fmt.Println(result)
	}
}