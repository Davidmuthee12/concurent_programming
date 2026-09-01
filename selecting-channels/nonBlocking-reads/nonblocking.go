package main

import (
	"fmt"
	"time"
)

func sendMsgAfter(seconds time.Duration) <-chan string{
	messages := make(chan  string)
	go func () {
		time.Sleep(seconds)
		messages <- "Hello"
	}()
	return messages
}

func main() {
	messages := sendMsgAfter(3 * time.Second) // sends channel message after 3 seconds
	for {
		select {
		case msg := <-messages: // Reads a message from the channel if theres one
			fmt.Println("Message Received", msg)
			return // when a message is available terminate the execution
		default: // when no message is availble the default case is executed
			fmt.Println("No messages waiting")
			time.Sleep(1 * time.Second)
		}
	}
}