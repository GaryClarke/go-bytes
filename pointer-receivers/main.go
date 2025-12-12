package main

import "fmt"

type Counter struct {
	Value int
}

func (c Counter) Increment() {
	c.Value++
}

type Account struct {
	Owner   string
	Balance float64
}

func (a *Account) Deposit(amount float64) {
	a.Balance += amount
}

func (a Account) BalanceInfo() string {
	return fmt.Sprintf("%s has a balance of %.2f", a.Owner, a.Balance)
}

func main() {
	acc := Account{Owner: "Alice", Balance: 100}
	acc.Deposit(50)
	acc.Deposit(25)
	fmt.Println(acc.BalanceInfo())
}
