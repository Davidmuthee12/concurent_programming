package main

import (
	"fmt"
	"sync"
	"time"
)

func playerHandler(cond *sync.Cond, playersRemaining *int, playerId int) {
	cond.L.Lock()
	fmt.Println(playerId, ": connected")
	*playersRemaining--
	if *playersRemaining == 0 {
		cond.Broadcast()
	}
	for *playersRemaining > 0 {
		fmt.Println(playerId, ":waiting for more players")
		cond.Wait()
	}
	cond.L.Unlock()
	fmt.Println("All players are connected. Ready Player: ", playerId)
	// Game Started
}

func main() {
	cond := sync.NewCond(&sync.Mutex{})
	playersINGame := 4

	for playerId := 0; playerId < 4; playerId++ {
		go playerHandler(cond, &playersINGame, playerId)
		time.Sleep(1 * time.Second)
	}
}