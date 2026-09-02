package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func sendMsgAfter(seconds time.Duration) <-chan string {
	// Sends the message "Hello" on the returned channel after the specified number of seconds
	messages := make(chan string)
	go func ()  {
		time.Sleep(seconds)
		messages <- "Hello"
	}()
	return messages
}

func main() {
	// Reads the timeout value from the program argument
	t, _ := strconv.Atoi(os.Args[1])
	// starts a goroutine that sends a message on the returned channel after 3 seconds
	messages := sendMsgAfter(3 * time.Second)
	timeoutDuration := time.Duration(t) * time.Second
	fmt.Printf("Waiting for message for %d seconds....\n", t)
	select {
	case msg := <-messages: // Read a message from the message channel if theres one
		fmt.Println("Message received:", msg)
	// creates a channel and timer, receiving a message after the specified duration
	case tNow := <-time.After(timeoutDuration):
		fmt.Println("Timed out. Waited untill:", tNow.Format("15:04:05"))
	}
}