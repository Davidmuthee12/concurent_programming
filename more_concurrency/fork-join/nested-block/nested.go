package main

import (
	"fmt"
	"math"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

type CodeDepth struct {
	file  string
	level int
}

// This program recursively scans through a directory and finds 
// the sourse file that has the deepest nested block

func deepestNestedBlock(filename string) CodeDepth {
	// Reads the full file into a menory buffer
	code, _ := os.ReadFile(filename)
	max := 0
	level := 0
	// iterate over every single character in the file
	for _, c := range code {
		if c == '{' {
			// when the character is an opening curly bracket,
			//  increments level by 1
			level += 1
			max = int(math.Max(float64(max), float64(level))) // Records the maximum value of level variable
		} else if c == '}' {
			// when the curly brackets are closed, decrements level by 1
			level -= 1
		}
	}
	// Returns the result with the filename
	return CodeDepth{filename, max} 
}

func forkIfNeeded(path string, info os.FileInfo, wg *sync.WaitGroup, results chan CodeDepth) {
	if !info.IsDir() && strings.HasSuffix(path, ".go") {
		// Adds 1 to the waitgroup
		wg.Add(1)
		// Spawns a new goroutine
		go func ()  {
			// calls the function and writes the
			//  return value on the common results channel
			results <- deepestNestedBlock(path)	
			wg.Done() // Marks the work done on the waitgroup
		}()
	}
}

func joinResults(partialResults chan CodeDepth) chan CodeDepth {
	finalResult := make(chan CodeDepth) // Create a channel that will contain the final result
	max := CodeDepth{"", 0}
	go func() {
		// Receives results from the channel until its closed
		for pr := range partialResults {
			// Records the value of the deepest nested block
			if pr.level > max.level {
				max = pr
			}
		}
		// After the channelis closed, writes result on the output channel
		finalResult <- max
	}()
	return finalResult
}

// We start by creating the common channel and the waitgroup
// Then we walk recursively through all the files in the directory specified in the arg
// and we folk a goroutine for each source file encountered
func main() {
	// Reads root directory from arguments
	dir := os.Args[1]
	// Creates common channel used by all forked goroutines
	partialResults := make(chan CodeDepth)
	wg := sync.WaitGroup{}

	// Walks the root directory, and for every file,
	//  calls the fork function, creating goroutine
	filepath.Walk(dir,
		func(path string, info os.FileInfo, err error) error {
			forkIfNeeded(path,info, &wg, partialResults)
			return nil
		})

	// calls the join function and gets the channel that will contain the final result
	finalResult := joinResults(partialResults)	

	// Waits for all the forked goroutines to complete their work
	wg.Wait()

	// Closes the common channel, signaling the join goroutine
	//  that the work is complete
	close(partialResults)

	// Receives the final result and outputs it on the console
	result := <-finalResult
	fmt.Printf("%s has the deepest nested code block of %d\n", result.file, result.level)
}