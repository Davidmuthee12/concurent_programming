package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func doWork(id int, wg *sync.WaitGroup) {
	i := rand.Intn(5)
	// sleeps for a random time
	time.Sleep(time.Duration(i) * time.Second)
	fmt.Println(id, "Done working after", i, "seconds")
	// signals that the goroutine has completed its tasks
	wg.Done()
}

func main() {
	wg := sync.WaitGroup{}
	// Adds 4 to the waitgroup since we have four pieces of work
	wg.Add(4)
	for i := 1; i <= 4; i++ {
		// Creates four goroutines, passing a reference to the waitgroup
		go doWork(i, &wg)
	}
	// waits for the work to be complete
	wg.Wait()
	fmt.Println("All complete")
}