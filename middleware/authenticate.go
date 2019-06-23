package middleware

import (
	"fmt"
	"go-server/database"
	"net/http"
)

func AuthenticateData(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// fmt.Println("excuting middlewareone")
		data := r.Header["Authorization"]
		secret := r.Header["Secret"]
		fmt.Println("secret", secret)
		// if err != nil {
		// 	fmt.Println("err", err)
		// }
		fmt.Println("data1111", data)
		fmt.Println("header", r.Header)
		result := database.IsSecretValid(map[string]string{"key": data[0], "secret": secret[0]})
		fmt.Println("result", result)
		next.ServeHTTP(w, r)
		// fmt.Println("excuting middlewaretwo")
	})
}
