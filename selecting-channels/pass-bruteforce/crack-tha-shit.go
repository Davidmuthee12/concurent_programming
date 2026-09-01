package main

import (
	"fmt"
	"time"
)

const (
	passwordToGuess = "go far"                      // set the password that we need to guess
	alphabet        = " abcdefghijklmnopqrstuvwxyz" // Define allpossible characters that the password is made of
)

func toBase27(n int) string {
	result := ""
	for n > 0 {
		// Algorithm convert a decimal integert into
		//  a string of base 27 using the alphabet constant
		result = string(alphabet[n%27]) + result
		n /= 27
	}
	return result
}

func guessPassword(from int, upto int, stop chan int, result chan string) {
	// loops over all password combinations using from and upto as starting and endpoints
	for guessN := from; guessN < upto; guessN += 1 {
		select {

		case <-stop:
			// Upon receiving a message on the stop channel,
			//  outputs a message and stops processing
			fmt.Printf("Stopped at %d [%d,%d]\n", guessN, from, upto)
			return

		default:
			// checks whether the password matches 
			// (in a real-life system we would try to protected resources)
			if toBase27(guessN) == passwordToGuess {
				result <- toBase27(guessN) // sends matching password on the result channel
				close(stop) // Closes the channelso that other goroutines stop checking the password
				return
			}
		}
	}
	fmt.Printf("Not found between [%d,%d]\n", from, upto)
}

func main() {
	// Creates a common channel used in the goroutine
	//  that signals when a password has been found
	finished := make(chan int)

	// Creates a channel that will contain the 
	// discovered password after its found
	passwordFound := make(chan string)

	for i := 1; i <= 387_420_488; i += 10_000_000 {
		// creates a goroutine with inpu ranges
		//  [1, 10m],[10m, 20m],....[380m,390m]
		go guessPassword(i, i+ 10_000_000, finished,passwordFound)
	}

	// waits for the password to be found
	fmt.Println("Password Found:", <-passwordFound)
	close(passwordFound)
	// Simulates the program using the password to access the resource
	time.Sleep(5 * time.Second)
}