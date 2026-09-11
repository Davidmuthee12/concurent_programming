package main

import (
	"fmt"
	"sync"
	"time"
)

func red(lock1, lock2 *sync.Mutex) {
	for {
		fmt.Println("Red: Acquiring Lock1")
		// Acquires and holds both locks
		lock1.Lock()
		fmt.Println("Red: Acquiring lock2")
		lock2.Lock()
		fmt.Println("Red: Both Locks Acquired")
		lock1.Unlock()
		lock2.Unlock() // Release both locks
		fmt.Println("Red: locks released")
	}
}

func blue(lock1, lock2 *sync.Mutex) {
	for {
		fmt.Println("Blue: Acquiring lock2")
		// Acquires and holds both locks
		lock2.Lock()
		fmt.Println("Blue: Acquiring Lock1")
		lock1.Lock()
		fmt.Println("Blue: Both locks Acquired")
		// Releases both locks
		lock1.Unlock()
		lock2.Unlock()
		fmt.Println("Blue: Locks Released")
	}
}

func main() {
	lockA := sync.Mutex{}
	lockB := sync.Mutex{}
	// Starts red() goroutine
	go red(&lockA, &lockB)
	// Starts blue() goroutine
	go blue(&lockA, &lockB)
	// Allows the red() and blue() goroutines torun for 20 seconds
	time.Sleep(20 * time.Second)
	fmt.Println("Done")
}