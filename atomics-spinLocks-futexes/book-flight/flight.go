package main

import (
	"fmt"
	"sort"
	"sync"
	"sync/atomic"
)

type Flight struct {
	Origin, Dest string
	SeatsLeft    int
	Locker       sync.Locker // Provides an interface containing lock and unlock functions
}

func Book(flights []*Flight, seatsToBook int) bool {
	bookable := true
	// Sorts flight in alphabetical order based on their origin and destination
	sort.Slice(flights, func(a, b int) bool {
		flightA := flights[a].Origin + flights[a].Dest
		flightB := flights[b].Origin + flights[b].Dest
		return  flightA < (flightB)
	})
	// Locks all the requested flights
	for _, f := range flights {
		f.Locker.Lock()
	}
	// Checks to see that all the requested flights have enough seats
	for i := 0; i < len(flights) && bookable; i++ {
		if flights[i].SeatsLeft < seatsToBook {
			bookable = false
		}
	}
	// Subtracts the seats from each flight only if there are enough seats for the entire booking
	for i := 0; i < len(flights) && bookable; i++ {
		flights[i].SeatsLeft-=seatsToBook
	}
	// Unlocks all the locked flights
	for _, f := range flights {
		f.Locker.Unlock()
	}
	return bookable // Returns the result of the booking
}

func main() {
	// Sets the variable to have the same value as the old parameter on CompareAndSwap()
	number := int32(17)
	result := atomic.CompareAndSwapInt32(&number, 17, 19) // Changes the value of the variable and returns true
	fmt.Printf("17 <- swap(17, 19): result %t, value: %d\n", result,number)
	// Sets the variable to have a different value than the older parameter on CompareAndSwap()
	number = int32(23)
	// Compares and fails, leaving the value of the variable unchanges, and returns false
	result = atomic.CompareAndSwapInt32(&number, 17, 19)
	fmt.Printf("23 <- swap(17, 19): result %t, value: %d\n", result, number)
}