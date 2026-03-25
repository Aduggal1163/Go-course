package main

import (
	"fmt"
	"time"
)

//	func main() {
//		firstname := getUserData("Please enter your first name:")
//		lastname := getUserData("Please enter your last name:")
//		birthdate := getUserData("Please enter your birthdate (MM/DD/YYYY):")
// 		outptUserDetails(firstname,lastname,birthdate)
//	}
// func outptUserDetails(firstname,lastname,birthdate) {
//		fmt.Println(firstname,lastname,birthdate)
// }
// func getUserData(promptText string) string {
// 	fmt.Print(promptText)
// 	var value string
// 	fmt.Scan(&value)
// 	return value
// }

type user struct {
	firstName string
	lastName  string
	birthdate string
	createdAt time.Time
}

// now this func belongs to that struct and now known as method (u user) is called reciever argument
func (u user) outptUserDetails() {
	fmt.Println("FirstName:", u.firstName, "\nLastName:", u.lastName, "\nBirthDate:", u.birthdate, "\nTime created:", u.createdAt)
}
func (u *user) clearUserName() {
	u.firstName = ""
	u.lastName = ""
}

func main() {
	userFirstname := getUserData("Please enter your first name: ")
	userLastname := getUserData("Please enter your last name: ")
	userBirthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")
	var appUser user
	//struct literal notation
	appUser = user{
		firstName: userFirstname,
		lastName:  userLastname,
		birthdate: userBirthdate,
		createdAt: time.Now(),
	}
	appUser.outptUserDetails()
	appUser.clearUserName()
	appUser.outptUserDetails()

	// null struct can also be create
	// appUser = user{} this will assign null values
}
func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scan(&value)
	return value
}
