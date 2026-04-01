package main

import "fmt"

type str string

func (name str) log() {
	fmt.Println(name)
}

func main() {
	var name str
	name = "Abhishek Duggal"
	name.log()
}
