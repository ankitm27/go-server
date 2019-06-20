package delivery

import (
	"fmt"
	user "go-server/useCase"
	"net/http"
	"reflect"

	"github.com/gorilla/mux"
)

func UserDelivery() *mux.Router {
	mux := mux.NewRouter()
	mux.Handle("/signup", http.HandlerFunc(user.SignUp))
	fmt.Println("mux", reflect.TypeOf(mux))
	// fmt.Println("")
	return mux
}
