package main

import (
	"bank/fileOps"
	"fmt"
)

const accountBalanceFile = "balance.txt"

func main() {
	fmt.Println("Welcome to my bank application")
	var accountBalance, err = fileOps.GetFloatFromFile(accountBalanceFile)
	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("----------------")
	}
	for {
		presentOptions()
		var choice int
		fmt.Print("Enter Your Choice:")
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
			fileOps.WriteFloatToFile(accountBalance, accountBalanceFile)
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
			fileOps.WriteFloatToFile(accountBalance, accountBalanceFile)
		case 4:
			fmt.Println("Good to see you")
			return
		default:
			fmt.Println("Invalid choice")
			fmt.Println("Thanks for choosing our bank")
		}
	}
}
