package main

import (
	"crypto/sha256"
	"io"
	"os"
)

func FHash(filepath string) []byte {
	file, _ := os.Open(filepath) // open file
	defer file.Close()

	// calculate the hash code using the crypto sha256 library
	sha := sha256.New()
	io.Copy(sha, file)

	return sha.Sum(nil)
}