package main

import (
	"fmt"
	"sync"
	"time"
)

func playerHandler(cond *sync.Cond, playersRemaining *int, playerId int, timedOut *bool) {
	cond.L.Lock()
	fmt.Println(playerId, ":connected")
	*playersRemaining--
	if *playersRemaining == 0 {
		cond.Broadcast()
	}
	for *playersRemaining > 0 && !*timedOut {
		fmt.Println(playerId, ": waiting for more players")
		cond.Wait()
	}
	cond.L.Unlock()
	if *timedOut {
		fmt.Println("Timed out, proceeding with ", *playersRemaining, "players")
	} else {
		fmt.Println("All players are connected, Ready player: ", playerId)
	}
}

func timeout(cond *sync.Cond, timedOut *bool) {
	time.Sleep(3 *time.Second)
	cond.L.Lock()
	*timedOut = true
	cond.Broadcast()
	cond.L.Unlock()
}
func main() {
	cond := sync.NewCond(&sync.Mutex{})
	timedOut := false
	go timeout(cond, &timedOut)
	playersInGame := 4
	for i := 0; i < 4; i++ {
		go playerHandler(cond, &playersInGame, i, &timedOut)
		time.Sleep(1 * time.Second)
	}
}