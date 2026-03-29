package main

import (
	"fmt"

	"example.com/structs/account"
)

func main() {
	accountNumber := inputText("Enter account number: ")
	holderName := inputText("Enter holder name: ")
	balance := inputTextF("Enter your initial balance: ")
	var appUser *account.Account
	appUser = account.New(accountNumber, holderName, balance)
	appUser.Deposit(1000)
	appUser.Display()
}
func inputText(text string) string {
	var i string
	fmt.Print(text)
	fmt.Scanln(&i)
	return i
}
func inputTextF(text string) float64 {
	var i float64
	fmt.Print(text)
	fmt.Scanln(&i)
	return i
}
