package user

import (
	"errors"
	"fmt"
)

type User struct {
	firstName string
	lastName  string
	dob       int64
}

func New(firstname string, lastname string, dob int64) *User {
	return &User{
		firstName: firstname,
		lastName:  lastname,
		dob:       dob,
	}
}

func (u *User) FullName() string {
	return u.firstName + " " + u.lastName
}

func (u *User) OutputDetails() {
	fmt.Println(u.firstName, " ", u.lastName, " ", u.dob)
}

func (u *User) ClearOutputDetails() {
	u.firstName = ""
	u.lastName = ""
}

func (u *User) CalculateAge() (int64, error) {
	// fmt.Println("Your Current Age is : ", 2026-u.dob)
	age := 2026 - u.dob
	if u.dob < 1990 {
		return 0, errors.New("Birth year is below 1990")
	}
	return age, nil
}
