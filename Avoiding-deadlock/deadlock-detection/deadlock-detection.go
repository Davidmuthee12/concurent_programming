package main

import (
	"fmt"
	"sync"
	"time"
)

func lockBoth(lock1, lock2 *sync.Mutex, wg *sync.WaitGroup) {
	for i := 0; i < 1000; i++ {
		// Locks and unlocks both mutexex
		lock1.Lock(); lock2.Lock()
		lock1.Unlock(); lock2.Unlock()
	}
	wg.Done() // Marks the waitgroup as done
}

func main() {
	lockA, lockB := sync.Mutex{}, sync.Mutex{}
	wg := sync.WaitGroup{}
	wg.Add(2)

	// Starts two goroutines, locking both mutexes at the same time
	go lockBoth(&lockA, &lockB, &wg)
	go lockBoth(&lockB, &lockA, &wg)
	// Creates a goroutine that waits on the waitgroup before outputting a message
	go func() {
		wg.Wait()
		fmt.Println("Done waiting on waitgroup")
	}()
	// waits for 30 seconds
	time.Sleep(30 * time.Second)
	// Outputs a message, and then the program terminates
	fmt.Println("Done")
}