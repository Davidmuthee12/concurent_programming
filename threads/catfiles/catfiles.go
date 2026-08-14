package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

func printFile(filename string) {
	content, err := os.ReadFile(filename)
	if err != nil{
		log.Fatal(err)
	}
	fmt.Println(string(content))
}

func main() {
	filenames := os.Args[1:]
	for _, filename := range filenames {
		go printFile(filename)
	}
	time.Sleep(2 * time.Second)
}