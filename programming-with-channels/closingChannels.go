package main

import "fmt"

func printNumbers(numbers <-chan int, quit chan int) {
	go func() {
		// Consumes 10 items from the numbers channel
		for i := 0; i < 10; i++ {
			fmt.Println(<-numbers)
		}
		close(quit) // close the quit channel
	}()
}

func main() {
	// Create the numbers and quit channels
	numbers := make(chan int)
	quit := make(chan int)
	printNumbers(numbers, quit) // calls the printNumbers() function, passing the channel
	next := 0
	for i := 1; ; i++ {
		next += 1 // generates the next triangular number
		select {
		case numbers <- next: // send the number on the numbers channels
		case <-quit:
			// when the quit channel is unblocked, 
			// outputs the message and terminates the execution
			fmt.Println("Quitting number generation")
			return
		}
	}
}