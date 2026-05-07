package main

import "fmt"

func main() {
	// var no1 int
	// var no2 int
	// fmt.Print("Enter first Number: ")
	// fmt.Scan(&no1)
	// fmt.Print("Enter second Number: ")
	// fmt.Scan(&no2)
	// sum := add(no1,no2)
	// fmt.Println("Sum is:",sum)
	// no1 := 11
	// res := oddchecker(no1)
	// fmt.Println(res)
	// name := "Abhishek Duggal"
	// fmt.Println(sizeof(name))
	// fmt.Println(fact(10))

	// sum, product := helper(10, 20)
	// fmt.Println(sum, product)

	// arr := []int{3, 4, 5, 7, 2, 4, 3, 2}
	// fmt.Println(max_min(arr))
	// var number int
	// fmt.Print("Enter number you want to check for prime: ")
	// fmt.Scan(&number)
	// res := prime(number)
	// fmt.Printf("The number is: %v", res)
	str := "abhishek"
	// fmt.Println(vowel(str))
	fmt.Println(reverse(str))
}
func reverse(str string) string {
	//  A rune in Go is a single Unicode code point (basically a “character”).
	runes := []rune(str)
	n := len(runes)
	for i := 0; i < n/2; i++ {
		runes[i], runes[n-i-1] = runes[n-i-1], runes[i]
	}
	return string(runes)
}
func vowel(str string) int {
	cnt := 0
	for _, ch := range str {
		if ch == 'a' || ch == 'e' || ch == 'i' || ch == 'o' || ch == 'u' {
			cnt++
		}
	}
	return cnt
}
func prime(number int) string {
	if number < 2 {
		return "Neither Prime nor composite"
	}
	for i := 2; i <= number; i++ {
		if number%i == 0 {
			return "Not Prime"
		}
	}
	return "Prime"
}
func max_min(arr []int) (int, int) {
	if len(arr) == 0 {
		return 0, 0
	}
	maxval := arr[0]
	minval := arr[0]
	for _, val := range arr {
		if val > maxval {
			maxval = val
		}
		if val < minval {
			minval = val
		}
	}
	return maxval, minval
}
func helper(a, b int) (int, int) {
	return a + b, a * b
}
func fact(no int) int {
	if no == 1 {
		return 1
	}
	return no * fact(no-1)
}
func sizeof(name string) int {
	return len(name)
}
func oddchecker(no1 int) bool {
	return no1%2 != 0
}
func add(no1, no2 int) int {
	return no1 + no2
}
