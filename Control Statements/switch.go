package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const accountBalanceFile = "balance.txt"

func writeBalanceToFile(balance float64) {
	balanceText := fmt.Sprint(balance)
	os.WriteFile(accountBalanceFile, []byte(balanceText), 0644)
}

func readBalanceFromFile() (float64, error) {
	// data,_ := os.ReadFile(accountBalanceFile)
	data, err := os.ReadFile(accountBalanceFile)
	if err != nil {
		return 1000, errors.New("Failed to read File")
	}
	balanceText := string(data)
	// balance,_:=strconv.ParseFloat(balanceText,64)
	balance, err := strconv.ParseFloat(balanceText, 64)
	if err != nil {
		return 1000, errors.New("Failed to parse stored balance File")
	}
	return balance, nil
}

func main() {
	fmt.Println("Welcome to my bank application")
	var choice int
	var accountBalance, err = readBalanceFromFile()
	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("----------------")
	}
	for {
		fmt.Print("What do you want to do?\n 1. Check Balance\n 2. Deposit Money\n 3. Withdraw Money \n 4. Exit\nEnter Your Choice:\n ")
		fmt.Scan(&choice)
		switch choice {
		case 1:
			fmt.Printf("Your balance is : $%.2f\n", accountBalance)
		case 2:
			fmt.Print("Enter amt you want to add: ")
			var depositAmt float64
			fmt.Scan(&depositAmt)
			if depositAmt <= 0 {
				fmt.Println("Invalid amt. Must be greater than 0.")
				continue
			}
			accountBalance += depositAmt
			fmt.Printf("Updated Balance is: $%.2f\n", accountBalance)
			writeBalanceToFile(accountBalance)
		case 3:
			fmt.Print("Enter amt you want to debit: ")
			var debitAmt float64
			fmt.Scan(&debitAmt)
			if debitAmt <= 0 {
				fmt.Println("Invalid amt. Must be greater than 0.")
				continue
			}
			if debitAmt > accountBalance {
				fmt.Println("Withdrawn amt is bigger")
				continue
			}
			accountBalance -= debitAmt
			fmt.Printf("Updated Balance is : $%.2f\n", accountBalance)
			writeBalanceToFile(accountBalance)
		case 4:
			fmt.Println("Good to see you")
			return
		default:
			fmt.Println("Invalid choice")
			fmt.Println("Thanks for choosing our bank")
		}
	}
}
