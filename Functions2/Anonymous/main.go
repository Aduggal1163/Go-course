package main

//when there is a case when we pass funcitons into functions or where a function returns a functions,
//we can use anonymous function which allows us to define a fnc just in time we need instead of in advance
//A function becomes a closure only if it captures variables from outer scope.
import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	numbers = transformNumbers(&numbers, func(number int) int {
		return number * 2
	})

	fmt.Println(numbers)

	triple := createTransformer(3)
	numbers = transformNumbers(&numbers, triple)
	fmt.Println(numbers)
}

func transformNumbers(numbers *[]int, transform func(int) int) []int {
	dNumbers := []int{}
	for _, val := range *numbers {
		dNumbers = append(dNumbers, transform(val))
	}
	return dNumbers
}

func createTransformer(factor int) func(int) int {
	return func(number int) int {
		return number * factor
	}
}
