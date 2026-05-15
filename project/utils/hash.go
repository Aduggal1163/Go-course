package utils

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	hashedpassword, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(hashedpassword), err

}

func ComparePassword(hp, p string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(p), []byte(hp))
	return err == nil
}
