package main

import (
	"fmt"

	waitgroupSemaphore "github.com/Davidmuthee12/threads/waitgroups-and-barriers/waitgroup-with-semaphore"
)

func doWork(id int, wg *waitgroupSemaphore.WaitGrp) {
	fmt.Println(id, "Done working")
	// when goroutine is complete, it calls Done() on the waitgroup
	wg.Done()
}

func main() {
	wg := waitgroupSemaphore.NewWaitGrp(4)
	for i := 1; i <= 4; i++ {
		// creates a goroutine, passing a reference to the waitgroup
		go doWork(i, wg)
	}
	// waits on the waitgroup for the work to be complete
	wg.Wait()
	fmt.Println("All done and dusted")
}