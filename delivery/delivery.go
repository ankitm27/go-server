package delivery

import (
	"fmt"
	"net/http"
	"reflect"

	user "go-server/useCase"

	"github.com/gorilla/mux"
)

func userDelivery() {
	mux := mux.NewRouter()
	mux.Handle("/signup", http.HandlerFunc(user.SignUp))
	fmt.Println("type of", reflect.TypeOf(mux))
}
