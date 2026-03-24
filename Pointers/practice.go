package main

import "fmt"

func practice() {
	a := 100
	p := &a
	fmt.Println("a:",a)
	fmt.Println("addr of x:",p)
	fmt.Println("value at address:",*p)
	modifyValue(&a)
	var no1 int
	var no2 int
	no1 =10
	no2 = 20
	swapTwoNumber(&no1,&no2)
	fmt.Println("no1 is:",no1,"\nno2 is:",no2)
	z :=1000
	y :=&z
	w :=&y
	fmt.Println("W is:",**w)
}
func modifyValue(a *int) {
	fmt.Println("Modified Value of a is:",*a-10)
}
func swapTwoNumber(a *int, b *int) {
	var temp int
	temp = *a
	*a=*b
	*b=temp
}
