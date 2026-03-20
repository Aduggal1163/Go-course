package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome to my bank application")
	var choice int
	var accountBalance float64 = 1000
	for {
		fmt.Print("What do you want to do?\n 1. Check Balance\n 2. Deposit Money\n 3. Withdraw Money \n 4. Exit\nEnter Your Choice:\n ")
		fmt.Scan(&choice)
		if choice == 1 {
			fmt.Printf("Your balance is : $%.2f", accountBalance)
			
		} else if choice == 2 {
			fmt.Print("Enter amt you want to add: ")
			var depositAmt float64
			fmt.Scan(&depositAmt)
			if depositAmt <= 0 {
				fmt.Println("Invalid amt. Must be greater than 0.")
				continue
			}
			accountBalance += depositAmt
			fmt.Printf("Updated Balance is: $%.2f\n", accountBalance)
		} else if choice == 3 {
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
		} else if choice == 4 {
			fmt.Println("Good to see you")
			break
		} else {
			fmt.Println("Invalid choice")
			break
		}
	}
	fmt.Println("Thanks for choosing our bank")
}
