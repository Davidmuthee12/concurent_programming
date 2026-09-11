package main

import (
	"fmt"
	"math/rand"
	"sort"
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

type Arbitrator struct {
	// Stores accounts with their availability status, either free or in use
	accountsInUse map[string]bool
	// Condition variable to be used to suspend
	//  goroutines if accounts are not available
	cond          *sync.Cond
}

func NewArbitrator() *Arbitrator {
	return &Arbitrator{
		accountsInUse: map[string]bool{},
		cond: sync.NewCond(&sync.Mutex{}),
	}
}

// Suspending executions to avoid deadlock
func (a *Arbitrator) LockAccounts(ids... string) {
	// locks mutex on condition variable
	a.cond.L.Lock()
	for allAvailable := false; !allAvailable; { // loops untill all accounts are free
		allAvailable = true
		for _, id := range ids {
			// is an account is in use, suspends the execution of the goroutine
			if a.accountsInUse[id] {
				allAvailable = false
				a.cond.Wait()
			}
		}
	}
	// Once all accounts are available, 
	// marks requested accounts as in use
	for _, id := range ids {
		a.accountsInUse[id] = true
	}
	// Unlocks the mutex on the condition variable
	a.cond.L.Unlock()  
}

// using broadcasts to resume goroutines
func (a *Arbitrator) UnlockAccounts(ids... string) {
	// locks mutex on the conditional variable
	a.cond.L.Lock() 
	// marks account as free
	for _, id := range ids {
		a.accountsInUse[id] = false
	}
	// Broadcast to resume any suspended goroutines
	a.cond.Broadcast()
	// Unlocks the mutex on the condition variable
	a.cond.L.Unlock()
}

func (src *BankAccount) Transfer(to *BankAccount, amount int, tellerId int, arb *Arbitrator) {
	// places the source and target accounts into a slice
	accounts := []*BankAccount{src, to}
	// Sorts the slice containing both accounts by their ID
	sort.Slice(accounts, func(a, b int) bool {
		return accounts[a].id < accounts[b].id
	})
	fmt.Printf("%d locking %s's account\n", tellerId, accounts[0].id)
	// Locks the account with the lower order by ID
	accounts[0].mutex.Lock()
	fmt.Printf("%d locking %s's account\n", tellerId, accounts[1].id)
	// Locks the account with the higher order by ID
	accounts[1].mutex.Lock()
	src.balance -= amount
	to.balance += amount
	// Unlocks both accounts
	to.mutex.Unlock()
	src.mutex.Unlock()
	fmt.Printf("%d Unocked %s and %s\n", tellerId, src.id, to.id)
}

// main() function using the arbitrator
func main() {
	accounts := []BankAccount{
		*NewBankAccount("Sam"),
		*NewBankAccount("Paul"),
		*NewBankAccount("Amy"),
		*NewBankAccount("Mia"),
	}
	total := len(accounts)
	// Creates a new arbitrator to be used in the transfers
	arb := NewArbitrator()
	for i := 0; i< 4; i++ {
		go func(tellerId int) {
			for i := 1; i < 1000; i++ {
				from, to := rand.Intn(total), rand.Intn(total)
				for from == to {
					to = rand.Intn(total)
				}
				accounts[from].Transfer(&accounts[to], 10, tellerId, arb)
			}
			fmt.Println(tellerId, "COMPLETE")
		}(i)
	}
	time.Sleep(60 * time.Second)
}