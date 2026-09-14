package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func stingy(money *int32) {
	for i := 0; i < 1000000; i++ {
		// Adds $10 atomically to the shared money variable
		atomic.AddInt32(money, 10)
	}
	fmt.Println("Stingy Done")
}

func spendy(money *int32) {
	for i := 0;i < 1000000; i++ {
		// Subtracts $10 atomically from the shared money variable
		atomic.AddInt32(money, -10)
	}
	fmt.Println("Spendy Done")
}

// main() function using atomic variables
func main() {
	// Creates a 32-bit integer with a value of 100
	money := int32(100)
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func ()  {
		stingy(&money)
		wg.Done()
	}()
	go func ()  {
		spendy(&money)
		wg.Done()
	}()
	wg.Wait() // waits on the waitgroup untill both goroutines are done
	fmt.Println("Money in account: ", atomic.LoadInt32(&money)) // Reads the value of the shared money variable and outputs it on the console
}