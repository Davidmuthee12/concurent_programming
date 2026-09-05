package main

import (
	"fmt"
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

func main() {
	// creates the quit channel
	quit := make(chan int)
	defer close(quit)
	// calls the function to start the goroutine returning urls on the result channel
	results := generateUrls(quit)
	for results := range results { // read all results from the channel
		fmt.Println(results) // print the results
	}
}