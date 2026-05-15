package models

import (
	"errors"

	"example.com/m/db"
	"example.com/m/utils"
)

type User struct {
	ID       int64
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (u *User) Save() error {
	query := `
		INSERT INTO users (email,password) VALUES (?,?)
	`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	hashedPassword, err := utils.HashPassword(u.Password)
	if err != nil {
		return err
	}
	result, err := stmt.Exec(u.Email, hashedPassword)

	if err != nil {
		return err
	}

	userId, err := result.LastInsertId()
	u.ID = userId
	return err

}

func (u *User) ValidateUserCredentials() error {
	query := "SELECT password FROM users WHERE email = ?"
	row := db.DB.QueryRow(query, u.Email)
	var reterivedPswrd string
	err := row.Scan(&reterivedPswrd)
	if err != nil {
		return errors.New("Credentials Invalid")
	}
	isPasswordMatched := utils.ComparePassword(u.Password, reterivedPswrd)
	if !isPasswordMatched {
		return errors.New("Credentials Invalid")
	}
	return nil
}
