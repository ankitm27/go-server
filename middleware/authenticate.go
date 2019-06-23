package middleware

import (
	"fmt"
	"net/http"
)

func AuthenticateData(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("excuting middlewareone")
		next.ServeHTTP(w, r)
		fmt.Println("excuting middlewaretwo")
	})
}
