package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func grepdir(path  string, fileInfo os.FileInfo, searchStr string) {
	fullpath := filepath.Join(path, fileInfo.Name())
	if !fileInfo.IsDir() {
		content, err := os.ReadFile(fullpath)
		if err != nil {
			log.Fatal(err)
		}
		if strings.Contains(string(content), searchStr) {
			fmt.Println(fullpath, "contains a match with ",searchStr)
		} else {
			fmt.Println(fullpath,"does NOT contain a match with", searchStr)
		}
	}
}

func main() {
	searchStr := os.Args[1]
	path := os.Args[2]
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}
	for _, fileInfo := range files {
		go grepdir(path, fileInfo, searchStr)
	}
	time.Sleep(2 * time.Second)

}