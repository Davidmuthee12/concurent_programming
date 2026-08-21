package main

import (
	"fmt"
	"os"
	"sync"
	"time"
)

func Stingy(money *int, cond *sync.Cond) {
	for i := 0; i < 1000000; i++ {
		cond.L.Lock()
		*money += 10
		cond.Signal()
		cond.L.Unlock()
	}
	fmt.Println("Stingy is Done")
}

func Spendy(money *int, cond *sync.Cond) {
	for i := 0; i < 200000; i++ {
		cond.L.Lock()
		for *money < 50 {
			cond.Wait()
		}
		*money -= 50
		if *money < 0 {
			fmt.Println("Money balance is negative")
			os.Exit(1)
		}
		cond.L.Unlock()
	}
	fmt.Println("Spendy is Done!!")
}

func main() {
	money := 100	
	mutex := sync.Mutex{}
	cond := sync.NewCond(&mutex)

	go Spendy(&money, cond)
	go Stingy(&money, cond)

	time.Sleep(2 * time.Second)
	mutex.Lock()
	fmt.Println("Money in the bank: ", money)
	mutex.Unlock()

}