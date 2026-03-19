package main

import (
	"fmt"
)

func main() {
	var revenue float64
	var expenses float64
	var taxRate float64
	fmt.Print("Enter revenue: ")
	fmt.Scan(&revenue)
	fmt.Print("Enter expenses: ")
	fmt.Scan(&expenses)
	fmt.Print("Enter tax ratio: ")
	fmt.Scan(&taxRate)
	EBT := revenue - expenses
	profit := EBT * (1 - taxRate/100)
	// fmt.Println("EBT is :", EBT)
	// fmt.Println("profit is: ", profit)
	ratio := EBT / profit
	// fmt.Println("Ratio is :", ratio)
	fmt.Printf("ebt is : %.2f \nprofit is : %.2f \nratio is : %.2f ", EBT, profit,ratio)
 
}
