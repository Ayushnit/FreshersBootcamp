package main

import (
	"fmt"
	"sync"
)

var balance int=500

func deposit(value int, wg *sync.WaitGroup,mutex *sync.Mutex) {
	mutex.Lock()
	fmt.Printf("Depositing %d to account with balance: %d\n", value, balance)
	balance += value
	mutex.Unlock()
	wg.Done()
}

func withdraw(value int, wg *sync.WaitGroup,mutex *sync.Mutex) {
	mutex.Lock()
	if value<= balance {
		fmt.Printf("Withdrawing %d from account with balance: %d\n", value, balance)
		balance -= value
	} else {
		fmt.Printf("Insufficient balance")
	}
	mutex.Unlock()
	wg.Done()
}
func main() {

	var wg sync.WaitGroup
	var m sync.Mutex
	wg.Add(2)
	go withdraw(200, &wg,&m)
	go deposit(300, &wg,&m)
	wg.Wait()

	fmt.Printf("New Balance %d\n", balance)
}

