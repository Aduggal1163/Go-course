package main

import "fmt"

type TransformFn func(int64) int64
// type anotherComplexFn func(int, []string, map[string][]int) ([]int, string)

func main() {
	numbers := []int64{1, 2, 3, 4}
	// numbers = doubleNumbers(&numbers)
	numbers = transformNumbers(&numbers, double)
	fmt.Println(numbers)

	numbers = transformNumbers(&numbers, triple)
	fmt.Println(numbers)
}

// func transformNumbers(numbers *[]int64, transform func(int64) int64) []int64 {
// OR
func transformNumbers(numbers *[]int64, transform TransformFn) []int64 {

	dNumbers := []int64{}
	for _, val := range *numbers {
		dNumbers = append(dNumbers, transform(val))
	}
	return dNumbers
}

func doubleNumbers(numbers *[]int64) []int64 {
	dNumbers := []int64{}
	// for idx,val := range *numbers {
	for _, val := range *numbers {
		// dNumbers = append(dNumbers, val*2)
		dNumbers = append(dNumbers, double(val))
	}
	return dNumbers
}

func double(number int64) int64 {
	return number * 2
}

func triple(number int64) int64 {
	return number * 3
}
