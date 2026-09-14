package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
	"sync"
	"sync/atomic"
)

const allLetters = "abcdefghijklmnopqrstuvwxyz"

func countLetters(url string, frequency []int32) {
	resp, _ := http.Get(url)
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		panic("Server returning error code: " + resp.Status)
	}
	// Reads the body of the web page
	body, _ := io.ReadAll(resp.Body)
	for _, b := range body { // iterates over every letter contained in the body of the document
		c := strings.ToLower(string(b))
		// checks tosee if the letter is part of the English alphabet
		cIndex := strings.Index(allLetters, c)
		if cIndex >= 0 {
			// Uses an atomic add operation to increment the letter
			atomic.AddInt32(&frequency[cIndex], 1)
		}
	}
	fmt.Println("Completed: ", url)
}

// main() function for atomic letter counter
func main() {
	wg := sync.WaitGroup{}
	wg.Add(31)
	// Creates a slice with size 26 of type 32-bit integers
	var frequency = make([]int32, 26)
	for i := 1000; i <= 1030; i++  {
		url := fmt.Sprintf("https://rfc-editor.org/rfc/rfc%d.txt", i)
		go func ()  {
			countLetters(url, frequency)
			wg.Done()
		}()
	}
	// waits untill all goroutines are complete
	wg.Wait()
	for i, c := range allLetters {
		// Loads the value of each count from the frequency slice and outputs them on the console
		fmt.Printf("%c-%d ", c, atomic.LoadInt32(&frequency[i]))
	}
}