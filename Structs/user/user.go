package user

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
	firstName string
	lastName  string
	birthdate string
	createdAt time.Time
}

// now this func belongs to that struct and now known as method (u user) is called reciever argument
func (u User) OutptUserDetails() {
	fmt.Println("FirstName:", u.firstName, "\nLastName:", u.lastName, "\nBirthDate:", u.birthdate, "\nTime created:", u.createdAt)
}
func (u *User) ClearUserName() {
	u.firstName = ""
	u.lastName = ""
}

// constructor
func New(firstName, lastName, birthdate string) (*User, error) {
	//validations
	if firstName == "" || lastName == "" || birthdate == "" {
		return nil, errors.New("Please provide mendatory details")
	}

	return &User{
		firstName: firstName,
		lastName:  lastName,
		birthdate: birthdate,
		createdAt: time.Now(),
	}, nil
}
