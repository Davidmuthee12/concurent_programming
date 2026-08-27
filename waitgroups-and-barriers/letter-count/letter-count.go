package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
	"sync"
	"time"
)

const AllLetters = "abcdefghijklmnopqrstuvwxyz"

func countLetters(url string, frequency []int, mutex *sync.Mutex) {
    resp, _ := http.Get(url)
    defer resp.Body.Close()
    if resp.StatusCode != 200 {
        panic("Server returning error code: " + resp.Status)
    }
    body, _ := io.ReadAll(resp.Body)
    mutex.Lock()
    for _, b := range body {
        c := strings.ToLower(string(b))
        cIndex := strings.Index(AllLetters, c)
        if cIndex >= 0 {
            frequency[cIndex] += 1
        }
    }
    mutex.Unlock()
    fmt.Println("Completed:", url, time.Now().Format("15:04:05"))
}



func main() {
	// creates a new waitgroup
	wg := sync.WaitGroup{}
	// Adds a delta of 31--one for each web pageto be downloaded concurrently
	wg.Add(31)
	mutex := sync.Mutex{}
	var frequency = make([]int, 26)
	for i := 1000; i <= 1030; i++ {
		url := fmt.Sprintf("https://rfc-editor.org/rfc%d.txt", i)
		// Creates a goroutine with an anonymous function
		go func ()  {
			countLetters(url, frequency, &mutex)
			// calls Done() after it finishes counting letters
			wg.Done()
		}()
	}
	// waits untill all goroutines are complete
	wg.Wait()
	mutex.Lock()
	for i, c := range AllLetters {
		fmt.Printf("%c-%d", c, frequency[i])
	}
	mutex.Unlock()
}