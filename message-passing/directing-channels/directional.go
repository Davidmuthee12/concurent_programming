package main

import (
	"fmt"
	"time"
)

func receiver(messages <-chan int) { // declares a receive-only channel
	for {
		msg := <-messages // Receives messages from a channel
		fmt.Println(time.Now().Format("15:04:05"), "Received:", msg)
	}
}

func sender(messages chan<- int) { // Declare a send-only channel
	for i := 1; ; i++ {
		fmt.Println(time.Now().Format("15:04:05"), "Sending", i)
		// Sends a message on the channel every second
		messages <- i
		time.Sleep(1 * time.Second)
	}
}

func main() {
	msgChannel := make(chan int)
	go receiver(msgChannel)
	go sender(msgChannel)
	time.Sleep(5 *time.Second)
}