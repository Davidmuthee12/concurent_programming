package main

import "fmt"

func main() {
	var ch chan string = nil // create a nil channel
	ch <- "message" // blocks execution as it tries to send message on the nil channel
	fmt.Println("This is never printed")
}