package main

import "fmt"

func main() {
	// Create a new channell of type string
	msgChannel := make(chan string)
	// Start a new goroutine with a reference to the channel
	go receiver(msgChannel)
	fmt.Println("Sending Hello...")
	// Sends three string messages over the channel
	msgChannel <- "Hello"
	fmt.Println("Sending There...")
	msgChannel <- "there"
	fmt.Println("Sending Wamunyoro...")
	msgChannel <- "Wamunyoro"
}

func receiver(messages chan string){
	msg := ""
	// Continues while the message received is not "Wamunyoro"
	for msg != "Wamunyoro" {
		msg = <-messages // Reads the next message from the channel
		fmt.Println("Received: ", msg) // outputs the message on the console
	}
}