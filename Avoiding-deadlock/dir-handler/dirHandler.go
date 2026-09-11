package main

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

func handleDirectories(dirs <-chan string, files chan<- string) {
	// Reads the full directory path from the input topic
	for fullpath := range dirs {
		fmt.Println("Reading all files from", fullpath)
		// Reads the content of a directory
		filesInDir, _ := os.ReadDir(fullpath)
		fmt.Printf("Pushing %d files from %s\n", len(filesInDir), fullpath)
		// Feeds each item of the directory contents onto the output topic
		for _, file := range filesInDir {
			// Starts new goroutine that sends each file to the files channel
			go func(fp string) {
				files <- fp
			}(filepath.Join(fullpath,file.Name()))
		}
	}
}

func handleFiles(files chan string, dirs chan string) {
	for path := range files { // Reads the full path of a file
		file, _ := os.Open(path)
		// Reads information about the file
		fileInfo, _ := file.Stat()
		// if the file is a directory, writes it to the output channel
		if fileInfo.IsDir() {
			fmt.Printf("Pushing %s directory\n", fileInfo.Name())
			dirs <- path
		} else {
			// if the file is not a directory, displays file information on the console
			fmt.Printf("File %s, size: %dMB, last modified: %s\n",
				fileInfo.Name(), fileInfo.Size() / (1024 * 1024),
				fileInfo.ModTime().Format("15:04:05"))
		}
	}
}

// main() function creating file and directory handlers
func main() {
	// Creates files and directory channels
	filesChannel := make(chan string)
	dirsChannel := make(chan string)
	// Starts file and directory handler goroutines
	go handleFiles(filesChannel, dirsChannel)
	go handleDirectories(dirsChannel, filesChannel)
	// Feeds the directory from the arguments to the directory channel
	dirsChannel <- os.Args[1]
	// sleeps for 60 seconds
	time.Sleep(60 * time.Second)
}