package middleware

import (
	"fmt"
	"net/http"
)

func AuthenticateData(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// fmt.Println("excuting middlewareone")
		data := r.Header["Authorization"]
		// if err != nil {
		// 	fmt.Println("err", err)
		// }
		fmt.Println("data1111", data)
		next.ServeHTTP(w, r)
		// fmt.Println("excuting middlewaretwo")
	})
}
