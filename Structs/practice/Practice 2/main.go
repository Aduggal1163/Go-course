package main

import (
	"fmt"
	"example.com/user"
)

func main() {
	manager := &user.UserManager{}

	for {
		fmt.Println("\n1. Add User")
		fmt.Println("2. Find User")
		fmt.Println("3. Delete User")
		fmt.Println("4. Show All Users")
		fmt.Println("5. Exit")

		var choice int
		fmt.Print("Enter choice: ")
		fmt.Scan(&choice)

		switch choice {

		case 1:
			name := inputText("Enter name: ")
			manager.AddUser(name)

		case 2:
			name := inputText("Enter name to find: ")
			u, found := manager.FindUserByName(name)
			if found {
				fmt.Println("Found:", u.Name)
			} else {
				fmt.Println("User not found")
			}

		case 3:
			name := inputText("Enter name to delete: ")
			if manager.DeleteUser(name) {
				fmt.Println("Deleted")
			} else {
				fmt.Println("User not found")
			}

		case 4:
			fmt.Println("All Users:")
			for _, u := range manager.Users {
				fmt.Println("-", u.Name)
			}

		case 5:
			return

		default:
			fmt.Println("Invalid choice")
		}
	}
}

func inputText(prompt string) string {
	fmt.Print(prompt)
	var text string
	fmt.Scan(&text)
	return text
}
