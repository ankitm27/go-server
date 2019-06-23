package middleware

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func CreateHash(password []byte) string {
	hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.MinCost)
	if err != nil {
		fmt.Println("There is some problem, Please try after some time", err)
	}
	return string(hashedPassword)
}

func ComparePassword(hashedPassword []byte, password []byte) bool {
	err := bcrypt.CompareHashAndPassword(hashedPassword, password)
	if err != nil {
		fmt.Println("There is some problem, Please try after some time", err)
		return false
	}
	return true
}
