package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/post"
	"example.com/user"
)

func main() {
	userFirstName := inputText("Enter your first name: ")
	userLastName := inputText("Enter your last name: ")
	userBirthDate := inputText2("Enter your birth year: ")
	var Appuser *user.User
	Appuser = user.New(userFirstName, userLastName, userBirthDate)
	// Appuser.OutputDetails()
	// age, err := Appuser.CalculateAge()
	// if err != nil {
	// 	fmt.Print(err)
	// } else {
	// 	fmt.Println("Your Current Age is : ", age)
	// }
	// Appuser.ClearOutputDetails()
	// Appuser.OutputDetails()
	////////////////////////////////////////
	postTitle := inputText("Enter your post title: ")
	postContent := inputText("Enter your post content: ")
	var postUser *post.Post
	postUser = post.NewPost(postTitle, postContent, *Appuser)
	postUser.OutputDetails()

}

func inputText(promptText string) string {
	fmt.Print(promptText)
	reader := bufio.NewReader(os.Stdin)
	line, _ := reader.ReadString('\n')
	line = strings.TrimSpace(line)
	return line
}

func inputText2(promptText string) int64 {
	fmt.Print(promptText)
	reader := bufio.NewReader(os.Stdin)
	line, _ := reader.ReadString('\n')
	line = strings.TrimSpace(line)

	var year int64
	fmt.Sscan(line, &year)
	return year
}
