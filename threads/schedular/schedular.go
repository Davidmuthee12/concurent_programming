package main

import (
	"fmt"
	"runtime"
)

func sayHello() {
	fmt.Println("hello")
}

func main() {
	go sayHello()
	runtime.Gosched()
	fmt.Println("Finished")
}
