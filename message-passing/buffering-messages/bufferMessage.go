package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// create a new channel with buffer capacity of 3 messages
	msgChannel := make(chan int, 3)
	// Creates a waitgroup with a size of 1
	wGroup := sync.WaitGroup{}
	wGroup.Add(1)

	// Starts the receiver goroutine with the buffered channel and waitgroup
	go receiver(msgChannel, &wGroup)
	for i := 1; i <= 6; i++ {
		size :=len(msgChannel) // Reads the number of messages in the buffered channel
		fmt.Printf("%s Sending: %d. Buffer size: %d/n",
			time.Now().Format("15:04:05"), i, size)
		msgChannel <- 1 // Sends six integer messages from 1 to 6
	}
	msgChannel <- -1 // Sends a message containing -1
	wGroup.Wait() // wait on the group until the receiver is finished
}

func receiver(messages chan int, wGroup *sync.WaitGroup) {
	msg := 0
	// Keeps messages in the channel until it receives a -1
	for msg != -1 {
		time.Sleep(1 * time.Second) // wait for 1 second
		msg = <-messages // Read the next message from the channel
		fmt.Println("Received: ", msg)
	}
	// Calls done() on the waitgroup after reading all the messages
	wGroup.Done()

}