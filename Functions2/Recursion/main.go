package main

import "fmt"

func main() {
	fact := factorial(5);
	fmt.Println(fact)
}
func factorial(num int) int {
	if num==1 {
		return 1
	}
	return num*factorial(num-1)
}