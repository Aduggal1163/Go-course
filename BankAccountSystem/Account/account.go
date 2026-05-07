package account

import (
	"errors"
	"fmt"
)

type Account struct {
	accountNumber string
	holderName    string
	balance       float64
}

type Bank struct {
	name    string
	account []*Account
}


func New(accountNumber, holderName string, balance float64) *Account {
	return &Account{
		accountNumber: accountNumber,
		holderName:    holderName,
		balance:       balance,
	}
}
func (a *Account) Display() {
	fmt.Println(a.accountNumber, a.holderName, "$", a.balance)
}
func (a *Account) Deposit(amount float64) error {
	if amount <= 0 {
		return errors.New("Amount must be greater then 0")
	}
	a.balance += amount
	return nil
}
func (a *Account) Withdraw(amount float64) error {
	if a.balance-amount <= 0 {
		return errors.New("withdrawn amnt is large")
	}
	a.balance -= amount
	return nil
}
