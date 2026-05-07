package main

import (
	"fmt"
)
func main() {
	var revenue float64
	var expenses float64
	var taxRate float64
	outputText("Enter revenue: ")
	fmt.Scan(&revenue)
	outputText("Enter expenses: ")
	fmt.Scan(&expenses)
	outputText("Enter tax ratio: ")
	fmt.Scan(&taxRate)
	EBT := calculateEBT(revenue, expenses)
	profit := calculateProfit(EBT, taxRate)
	ratio := EBT / profit
	fmt.Printf("ebt is : %.2f \nprofit is : %.2f \nratio is : %.2f ", EBT, profit, ratio)
}
func outputText(title string) {
	fmt.Print(title)
}
func calculateEBT(revenue float64, expenses float64) float64 {
	return revenue - expenses
}
func calculateProfit(EBT float64, taxRate float64) (sol float64) {
	sol = EBT * (1 - taxRate/100)
	
	return
}
