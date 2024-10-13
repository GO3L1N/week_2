package main

import (
	"fmt"
)

type BankAccount struct {
	holderName string
	balance    float64
}

func (account *BankAccount) Deposit(amount float64) {
	if amount > 0 {
		account.balance += amount
		fmt.Printf("Пополнено: %.2f. Текущий баланс: %.2f.\n", amount, account.balance)
	} else {
		fmt.Println("Сумма пополнения должна быть положительной.")
	}
}

func (account *BankAccount) Withdraw(amount float64) {
	if amount > 0 && amount <= account.balance {
		account.balance -= amount
		fmt.Printf("Снято: %.2f. Текущий баланс: %.2f.\n", amount, account.balance)
	} else if amount > account.balance {
		fmt.Printf("Недостаточно средств для снятия %.2f. Текущий баланс: %.2f.\n", amount, account.balance)
	} else {
		fmt.Println("Сумма снятия должна быть положительной.")
	}
}

func main() {
	account := BankAccount{holderName: "Иван Иванов", balance: 1000.00}

	fmt.Printf("Владелец счета: %s, Начальный баланс: %.2f\n", account.holderName, account.balance)

	account.Deposit(500.00)

	account.Withdraw(200.00)

	account.Withdraw(1500.00)

	fmt.Printf("Финальный баланс: %.2f\n", account.balance)
}
