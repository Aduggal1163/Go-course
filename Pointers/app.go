package main

import "fmt"

func main() {
	x := 5
	y := 30
	changeValue(&x, y)
	fmt.Println("X:", x)
	fmt.Println("y:", y)
	var agepointer *int
	agepointer = &y
	fmt.Println("age pointer is:", *agepointer)
	fmt.Println("Age above 18 is:", getAdultAge(agepointer))
	practice()
}
func changeValue(x *int, y int) {
	*x = *x + y
	y = y + 10
}
func getAdultAge(age *int) int {
	return *age - 18
}
