package main

import (
	"fmt"
	"time"
)

func doWork(id int) {
	fmt.Printf("work %d started at %s\n",id, time.Now().Format("15:04:05"))
	time.Sleep(1 * time.Second)
	fmt.Printf("work %d finisshed at %s\n", id, time.Now().Format("15:04:05"))
}

func main() {
	for i := 0; i < 5; i++ {
		doWork(i)
	}
}