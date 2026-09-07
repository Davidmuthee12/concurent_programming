package main

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"
)

//  This function computes the hash code
//  of each and every file in a directory
func main() {
	dir := os.Args[1]

	files, _ := os.ReadDir(dir) // Gets a list of files from the specified directory
	wg := sync.WaitGroup{}
	for _, file := range files {
		if !file.IsDir() {
			wg.Add(1)
			// Starts a goroutine to compute the hash code for the file in the itaration
			go func(filename string) {
				fPath := filepath.Join(dir, filename)
				// Computes and outputs the hash code for the file using the previously developed function
				hash := FHash(fPath)
				fmt.Printf("%s - %x\n", filename, hash)
				wg.Done()
			}(file.Name())
		}
	}
	// Waits for all the tasks, which are computing the hashes, to be complete
	wg.Wait()
}