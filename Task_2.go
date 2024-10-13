package main

import "fmt"

type Account struct {
	balance float64
}

func (a *Account) GetBalance() float64 {
	return a.balance
}

func (a *Account) Deposit(amount float64) {
	a.balance += amount
	fmt.Println("Аккаунт пополнен: ", a.balance)
}

func (a *Account) Withdraw(amount2 float64) {
	a.balance -= amount2
	fmt.Println("Аккаунт уменьшился: ", a.balance)
}

func main() {
	acc := &Account{1000.00}
	acc.Deposit(52.88)
	acc.Withdraw(2.88)
	fmt.Printf("Текущий баланс: %.2f\n", acc.GetBalance())
}
