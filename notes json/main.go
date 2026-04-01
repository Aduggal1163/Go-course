package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/note/note"
)

func main() {
	title := getUserInput("Note Title: ")
	content := getUserInput("Content Title:")
	userNote, err := note.New(title, content)
	if err != nil {
		fmt.Println(err)
		return
	}
	userNote.Display()
	err = userNote.Save()
	if err != nil {
		fmt.Println("\nSAVING FAILED")
		return
	}
	fmt.Println("\nSAVING SUCCESSFUL")
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
