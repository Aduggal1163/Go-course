package main

// import (
// 	"fmt"
// 	"math"
// )

// func main() {
// 	var investmentAmount float64= 1000
// 	var expectedReturnRate = 5.5
// 	var years float64= 10
// 	var futureValue = investmentAmount * math.Pow(1+expectedReturnRate/100,years)
// 	fmt.Println("The future value is :", futureValue)
// }

//  /// //// // //  OR /// // // / // // / /

// func main() {
// 	var investmentAmount = 1000
// 	var expectedReturnRate = 5.5
// 	var years = 10
// 	var futureValue = float64(investmentAmount) * math.Pow(1+expectedReturnRate/100, float64(years))
// 	fmt.Println("The future value is : ", futureValue)
// }

////////// OR ///////////

// func main() {
// 	const inflationRate = 2.5
// 	var investmentAmount = 1000
// 	expectedReturnRate := 5.5
// 	// this is shortcut and recommended way to declare and assign a variable with this now u cant do like this expectedReturnRate float64 := 5.5
// 	var years float64 = 10
// 	futureValue := float64(investmentAmount) * math.Pow(1+expectedReturnRate/100, years)
// 	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)
// 	fmt.Println("Future Value is :",futureValue)
// 	fmt.Println(" inflation adjusted value is : ",futureRealValue)

// }

////////// OR /////////////

// func main() {
// const inflationRate = 2.5
// var investmentAmount float64
// var years  float64// if we dont initiliaze the value that type must be declared too
// var expectedReturnRate float64
// fmt.Print("Investment Amount :")
// fmt.Scan(&investmentAmount)
// fmt.Print("Years: ")
// fmt.Scan(&years)
// fmt.Print("Expected Return Rate: ")
// fmt.Scan(&expectedReturnRate)

// futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100,years)
// futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

// fmt.Printf("Future value is : %v \n",futureValue);

// fmt.Printf("Future inflated amount is: %.2f \n",futureRealValue);

// }
