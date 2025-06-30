package main

import (
	"errors"
	"fmt"
)

// СТРУКТУРА: Банковский аккаунт
type BankAccount struct {
	Owner   string
	Balance float64
}

// ФУНКЦИЯ: пополнение счёта
func (acc *BankAccount) Deposit(amount float64) {
	acc.Balance += amount
}

// ФУНКЦИЯ: снятие со счёта с проверкой баланса
func (acc *BankAccount) Withdraw(amount float64) error {
	if amount > acc.Balance {
		return errors.New("недостаточно средств")
	}
	acc.Balance -= amount
	return nil
}

// ФУНКЦИЯ: получение текущего баланса
func (acc BankAccount) GetBalance() float64 {
	return acc.Balance
}

func main() {
	account := BankAccount{Owner: "Антон", Balance: 1000}

	account.Deposit(500)
	fmt.Printf("Баланс после пополнения: %.2f\n", account.GetBalance())

	err := account.Withdraw(2000)
	if err != nil {
		fmt.Println("Ошибка при снятии:", err)
	} else {
		fmt.Printf("Баланс после снятия: %.2f\n", account.GetBalance())
	}

	err = account.Withdraw(300)
	if err != nil {
		fmt.Println("Ошибка при снятии:", err)
	} else {
		fmt.Printf("Баланс после снятия: %.2f\n", account.GetBalance())
	}
}
