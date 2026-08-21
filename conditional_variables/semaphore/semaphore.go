package main

import (
	"fmt"
	"sync"
)

type semaphore struct {
	// permits remaining on the semaphore
	permits int
	// condition variable used for waiting
	// when there are not enough permits
	cond    *sync.Cond
}

func NewSemaphore(n int) *semaphore {
	return &semaphore{
		permits: n, // initial number or permits on the new semaphore
		cond: sync.NewCond(&sync.Mutex{}), // here we initialize a new condition variable and associated mutex on the new semaphore
	}
}

func (rw *semaphore) Acquire() {
	// Acquires mutex to protect permits variable
	rw.cond.L.Lock()
	for rw.permits <= 0 {
		// wait until there is an available permit
		rw.cond.Wait()
	}
	// Decrease the number of permits by 1
	rw.permits-- 
	rw.cond.L.Unlock()
}

func (rw *semaphore) Release() {
	// Acquire mutex to protect permits variable
	rw.cond.L.Lock()

	// increases the number of availbale permits by 1
	rw.permits++
	// signal condition variable that one more permit is available
	rw.cond.Signal()

	rw.cond.L.Unlock()
}

func doWork(semaphore *semaphore) {
	fmt.Println("work started")
	fmt.Println("work finished")
	// when the goroutine finishes, it releases a permit to notify the main() goroutine
	semaphore.Release()
}

func main() {
	// Create a new semaphore using the previous implementation
	semaphore := NewSemaphore(0)
	for i := 0; i < 50000; i++ {
		// start goroutine passing a reference to the semaphore
		go doWork(semaphore)
		fmt.Println("waiting for child goroutine...")
		// wait for an available permit on the semaphore indicating the task is complete
		semaphore.Acquire()
		fmt.Println("Child goroutine finished")
	}
}