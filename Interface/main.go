package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/note/note"
	"example.com/note/todo"
)

type saver interface {
	Save() error
}
type displayer interface {
	Display()
}

// embedded interface
type outputtable interface {
	saver
	displayer
	// doSomething(int) string
}

// type outputtable interface {
// 	Save() error
// 	Display()
// }

// common convention of making interface is that if your interface has 1 method then the interface name is method name + er
// func printSomething(value interface{})
// OR
func printSomething(value any) {
	// typedVal,ok  := value.(int)
	// if !ok {
	// 	fmt.Println("Value not int")
	// }
	// fmt.Println("Integer:",typedVal)

	switch value.(type) {
	case int:
		fmt.Println("Integer:", value)
	case float64:
		fmt.Println("Float:", value)
	case string:
		fmt.Println("String:", value)
	default:
		fmt.Println("Type not available")
	}
}

func main() {
	printSomething(1)
	// printSomething(1.3)
	// printSomething("Abhi")
	title := getUserInput("Note Title: ")
	content := getUserInput("Note Content: ")
	text := getUserInput("Todo Text: ")

	todo, err := todo.New(text)
	if err != nil {
		fmt.Println(err)
		return
	}
	userNote, err := note.New(title, content)
	if err != nil {
		fmt.Println(err)
		return
	}

	outputData(todo, "Todo")
	outputData(userNote, "Note")
}

// We can see duplicate code for saving todo and note.
// This can be resolved by creating a saveData function that handles both.
// While this could be done with a normal function, it would not be generic or clean.
// Using an interface provides a scalable and maintainable solution.

func outputData(data outputtable, label string) {
	data.Display()
	err := saveData(data, label)
	if err != nil {
		return
	}
}

func saveData(data saver, label string) error {
	err := data.Save()
	if err != nil {
		fmt.Println("\nSAVING", label, "FAILED")
		return err
	}
	fmt.Println("\nSAVING", label, "SUCCESSFUL")
	return nil
}

func getUserInput(promptText string) string {
	fmt.Print(promptText)
	reader := bufio.NewReader(os.Stdin)  //read user input for larger strings including space
	text, err := reader.ReadString('\n') //where to stop reading
	if err != nil {
		return ""
	}
	text = strings.TrimSpace(text)
	return text
}
