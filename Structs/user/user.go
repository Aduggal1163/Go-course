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

// constructor embedding (Inheritance)
type Admin struct {
	email    string
	password string
	User
}

// now this func belongs to that struct and now known as method (u user) is called reciever argument
func (u User) OutptUserDetails() {
	fmt.Println("FirstName:", u.firstName, "\nLastName:", u.lastName, "\nBirthDate:", u.birthdate, "\nTime created:", u.createdAt)
}

// koi bhi field ko modilfy krne layi phir hum ko use karna pdega pointer
func (u *User) ClearUserName() {
	u.firstName = ""
	u.lastName = ""
}

// constructor
func NewAdmin(email, password, firstName, lastName, birthdate string) *Admin {
	return &Admin{
		email:    email,
		password: password,
		User: User{
			firstName: "admin",
			lastName:  "admin",
			birthdate: birthdate,
			createdAt: time.Now(),
		},
	}
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