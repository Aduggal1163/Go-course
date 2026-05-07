package main

import (
	"fmt"
	"example.com/structs/user"
)

func main() {
	userFirstname := getUserData("Please enter your first name: ")
	userLastname := getUserData("Please enter your last name: ")
	userBirthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")
	var appUser *user.User
	appUser, err := user.New(userFirstname, userLastname, userBirthdate)
	if err != nil {
		fmt.Println("User creation fail", err)
		return
	}
	admin := user.NewAdmin("textexample.com", "test123", userFirstname, userLastname, userBirthdate)
	admin.OutptUserDetails()
	admin.ClearUserName()
	admin.OutptUserDetails()

	//struct literal notation
	// appUser = user{
	// 	firstName: userFirstname,
	// 	lastName:  userLastname,
	// 	birthdate: userBirthdate,
	// 	createdAt: time.Now(),
	// }
	
	appUser.OutptUserDetails()
	appUser.ClearUserName()
	appUser.OutptUserDetails()

	// null struct can also be create
	// appUser = user{} this will assign null values
}
func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	// we use Scanln instead of Scan because it allows us to if we press enter it will go into next input not waiting for the current input
	return value
}

