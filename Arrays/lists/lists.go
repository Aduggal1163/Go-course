package lists

import (
	"fmt"
	"time"
)

type Product struct {
	title     string
	id        int32
	price     float64
	createdAt time.Time
}

func main() {
		// prices := [4]float64{10.0, 11.09, 12, 13.99}
	// 	fmt.Println(prices)

	// 	var productNames [4]string
	// 	productNames[0] = "Clg"
	// 	productNames[1] = "Gym"
	// 	productNames[2] = "Home"
	// 	productNames[3] = "Work"
	// 	fmt.Println(productNames)

	// 	var arr [3]int
	// 	for i := 0; i < 3; i++ {
	// 		fmt.Scan(&arr[i])
	// 	}
	// 	fmt.Println(arr)

	// 	featuredProductNames := productNames[2:4]
	// 	fmt.Println(featuredProductNames)

	// 	price := prices[:2]
	// 	price2 := prices[1:]
	// 	highlightedPrice := price[1:]
	// 	fmt.Println(price, "and highlighted price is: ", highlightedPrice)
	// 	fmt.Println("Length of sliced array is:", len(price), "capacityTowardsStart:", cap(price), "capacityTowardsEnd", cap(price2))

	// 	dynamic array
	// 	dynamicArray := []int{}
	// 	dynamicArray[0] = 1
	// 	dynamicArray[1] = 2
	// 	dynamicArray[2] = 3
	// 	dynamicArray[3] = 4
	// 	fmt.Println(dynamicArray)
	// 	this would cause error
	// 	dynamicArray = append(dynamicArray, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	// 	fmt.Println(dynamicArray)

	// PRACTICE MORE

	//1
	hobbies := [3]string{}
	hobbies[0] = "Cricket"
	hobbies[1] = "Coding"
	hobbies[2] = "Relaxing"
	fmt.Println(hobbies, hobbies[0])
	//2
	hobbiesLastSlice := hobbies[1:3]
	fmt.Println(hobbiesLastSlice)
	//3
	// hobbiesFirstSlice := hobbies[0:2]
	hobbiesFirstSlice := hobbies[:2]
	fmt.Println(hobbiesFirstSlice)
	//4
	hobbiesFirstUpdatedSlice := hobbiesFirstSlice[1:]
	hobbiesFirstUpdatedSlice = append(hobbiesFirstUpdatedSlice, hobbies[2])
	fmt.Println(hobbiesFirstUpdatedSlice)
	//5
	goals := []string{}
	goals = append(goals, "Software Developer", "CEO")
	fmt.Println(goals)
	//6
	goals[1] = "Chief Executive Officer"
	goals = append(goals, "Tribal Chief")
	fmt.Println(goals)
	//7
	products := []Product{}
	products = append(products,
		Product{"mouse", 101, 999.99, time.Now()},
		Product{"keyboard", 102, 1999.95, time.Now()},
	)
	
	fmt.Println(products)

	pricess := [] float64{10.0, 11.09, 12, 13.99}
	discountprices := [] float64{134,1223,232,23,2123}
	pricess = append(pricess,discountprices...)
	fmt.Println(pricess)

}
