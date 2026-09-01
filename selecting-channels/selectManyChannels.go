package main

import (
	"fmt"
	"time"
)

func writeEvery(msg string, seconds time.Duration) <-chan string {
	messages := make(chan string) // Create a channel of type string
	// Create a new anonymous function
	go func() {
		for {
			time.Sleep(seconds) // sleep for a specified period
			messages <- msg // sends the specified message on  the channel
		}
	}()
	return messages // returns the newly created messages on the channel
}
 
func main() {
	// Creates a goroutine sending messages every second on Channel A
	messagesFromA := writeEvery("Tick", 1 * time.Second)
	// Creates a goroutine sending messages every 3 seconds on Channel B
	messagesFromB := writeEvery("Tock", 3 * time.Second)

	for { // loops forever
		select {
		case msg1 := <-messagesFromA:
			fmt.Println(msg1) // outputs messages from channel A if one is available
		case msg2 := <-messagesFromB:
			fmt.Println(msg2) // outputs messages from channel B if one is available
		}
	}
}