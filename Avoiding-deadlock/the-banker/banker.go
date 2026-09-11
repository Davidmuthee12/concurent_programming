package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type BankAccount struct {
	id      string
	balance int
	mutex   sync.Mutex
}

func NewBankAccount(id string) *BankAccount {
	// Creates a new instance of a bank account with $100 and a new mutex
	return &BankAccount{
		id: id,
		balance: 100,
		mutex: sync.Mutex{},
	}
}

func (src *BankAccount) Transfer(to *BankAccount, amount int, exid int) {
	fmt.Printf("%d locking %s's account\n", exid, src.id)
	src.mutex.Lock() // locks mutex on the source account
	fmt.Printf("%d locking %s's account\n", exid, to.id)
	to.mutex.Lock() // locks mutex on the target account
	// subtract money from the source and adds it to the target account
	src.balance -= amount
	to.balance += amount
	// Unlocks both target and source accounts
	to.mutex.Unlock()
	src.mutex.Unlock()
	fmt.Printf("%d Unlocked %s and %s\n", exid, src.id, to.id)
}

func main() {
	accounts := []BankAccount{
		*NewBankAccount("sam"),
		*NewBankAccount("Paul"),
		*NewBankAccount("Amy"),
		*NewBankAccount("Mia khalifa"),
	}
	total := len(accounts)
	for i := 0;i < 4; i++ {
		go func(eId int) {
		// creates a goroutine with a unique execution ID
			for j := 1; j < 1000; j++ {
				// Selects a source and a target account for the transfer
				from, to := rand.Intn(total), rand.Intn(total)
				for from == to {
					to = rand.Intn(total)
				}
				accounts[from].Transfer(&accounts[to], 10, eId) // performs the transfer
			}
			// Once all the 1000 transfers are completed,
			// outputs the complete message
			fmt.Println(eId, "COMPLETE")
		}(i)
	}
	time.Sleep(60 * time.Second) // Waits 60 seconds before terminating the program
}