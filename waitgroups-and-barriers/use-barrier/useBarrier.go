package main

import (
	"fmt"
	"time"

	"github.com/Davidmuthee12/threads/waitgroups-and-barriers/barriers"
)

func workAndWait(name string, timeToWork int, barrier *barriers.Barrier) {
	start := time.Now()
	for {
		fmt.Println(time.Since(start), name, "is running")
		// simulate doing work for a number of seconds
		time.Sleep(time.Duration(timeToWork) * time.Second)
		fmt.Println(time.Since(start), name, "is waiting on barrier")
		// Waits for other goroutines to catch up
		barrier.Wait()
	}
}

func main() {
	// Creates a new barrier with two participants
	barrier := barriers.NewBarrier(2)

	go workAndWait("Red", 4, barrier)

	go workAndWait("Blue", 10, barrier)

	time.Sleep(100 * time.Second)

}