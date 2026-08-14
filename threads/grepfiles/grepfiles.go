package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func grepfile(filename string, searchStr string) {
	content, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	if strings.Contains(string(content), searchStr) {
		fmt.Println(filename, "contains a match with ", searchStr)
	} else {
		fmt.Println(filename, "does NOT sincerely contain a match of ", searchStr)
	}
}

func main() {
	searchStr := os.Args[1]
	filenames := os.Args[2:]
	for _, filename := range filenames {
		go grepfile(filename, searchStr)
	}
	time.Sleep(2 * time.Second)
}