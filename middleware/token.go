package middleware

import (
	"fmt"

	"github.com/dgrijalva/jwt-go"
)

func CreateToken(email string, password string) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": email,
		"password": password,
	})
	// fmt.Println("token111", token)
	tokenString, err := token.SignedString([]byte("secret"))
	if err != nil {
		fmt.Println("err", err)
	}
	// fmt.Println("token string11111", tokenString)
	return tokenString
}

func ValidateToken(tokenStr string) {
	token, _ := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("There was an error")
		}
		return []byte("secret"), nil
	})
	// fmt.Println("token111", token)
	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// type User struct {
		// username string
		// password string
		// }
		// var user User
		// data := mapstructure.Decode(claims, &user)
		// json.NewEncoder(w).Encode(user)
		// fmt.Println("claims", claims1)
		// fmt.Println("claims", claims1["username"])
		// fmt.Println("claims 1212", claims1["password"])
	} else {
		// json.NewEncoder(w).Encode(Exception{Message: "Invalid authorization token"})
	}
}
