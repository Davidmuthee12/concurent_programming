package main

import (
	"crypto/sha256"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func FHash(filepath string) []byte {
	file, _ := os.Open(filepath) // open file
	defer file.Close()

	// calculate the hash code using the crypto sha256 library
	sha := sha256.New()
	io.Copy(sha, file)

	return sha.Sum(nil)
}

// this function creates a single hash code for the entire directory
func main() {
	dir := os.Args[1]
	files, _ := os.ReadDir(dir)
	sha := sha256.New()
	var prev, next chan int
	for _, file := range files {
		if !file.IsDir() {
			// Creates the next channel used by the goroutine to signal that its ready
			next = make(chan int)
			go func(filename string, prev, next chan int) {
				fpath := filepath.Join(dir, filename)
				// Computes the hash code on the file
				hashOnFile := FHash(fpath)
				// if the goroutine is not in the first iteration, 
				// waits untill the previos iteration sends signal
				if prev != nil {
					<-prev
				}
				sha.Write(hashOnFile) // computes the directory partial hash
				next <- 0 // Signals tothe next iteration that its done
			}(file.Name(), prev, next)
			// Assigns the next channel to be previous;
			//  the goroutine will wait on a signal from the current iteration
			prev = next
		}
	}
	<-next // waits for the last iteration to be complete before outputting the result
	fmt.Printf("%x\n", sha.Sum(nil))
}