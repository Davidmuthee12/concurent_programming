package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

func filesearch(dir string, filename string, wg *sync.WaitGroup) {
	// Reads all files from the directory given to the function
	files, _ := os.ReadDir(dir)
	for _, file := range files {
		// Joins each file to the directory: e.g 'cat.jpg' becomes '/home/pics/cat.jpg'
		fpath := filepath.Join(dir, file.Name())
		if strings.Contains(file.Name(), filename) {
			// If there is a match, prints path on console
			fmt.Println(fpath)
		} else {
			fmt.Println("No such file found")
		}
		if file.IsDir() {
			// if it is a directory, adds 1 to the waitgroup before starting a new goroutine
			wg.Add(1)
			// creates goroutine recursively
			go filesearch(fpath,filename, wg)
		}

	}
	// Marks done on the waitgroup after processing all files
	wg.Done()
}

func main() {
	// Creates a new empty waitgroup
	wg := sync.WaitGroup{}
	// Adds a delta of 1 to the waitgroup
	wg.Add(1)
	// Creates a new goroutine, performing the file search and passing a refernec to the waitgroup
	go filesearch(os.Args[1], os.Args[2], &wg)
	// waits for the search to complete
	wg.Wait()
}