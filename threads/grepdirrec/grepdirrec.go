package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func grepdirrec(path string, dirEntry os.DirEntry, searchStr string) {
	fullpath := filepath.Join(path, dirEntry.Name())
	if dirEntry.IsDir(){
		files, err := os.ReadDir(fullpath)
		if err != nil {
			log.Fatal(err)
		}
		for _, file := range files {
			go grepdirrec(fullpath, file, searchStr)
		}
	} else {
		content, err := os.ReadFile(fullpath)
		if err != nil {
			log.Fatal(err)
		}
		if strings.Contains(string(content), searchStr) {
			fmt.Println(fullpath, "contains a match with ", searchStr)
		} else {
			fmt.Println(fullpath, "does NOT contain a match with ", searchStr)
		}
	}
}

func main() {
	searchStr := os.Args[1]
	path := os.Args[2]
	files, err := os.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}
	for _, dirEntry := range files {
		go grepdirrec(path, dirEntry, searchStr)
	}
	time.Sleep(2 * time.Second)
}