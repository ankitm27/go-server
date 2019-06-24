package delivery

import (
	user "go-server/useCase"
	"net/http"

	middleware "go-server/middleware"

	"github.com/gorilla/mux"
)

func UserDelivery() *mux.Router {
	mux := mux.NewRouter()
	mux.Handle("/signup", http.HandlerFunc(user.SignUp))
	mux.Handle("/getdatatypewise", middleware.AuthenticateData(http.HandlerFunc(user.GetDataTypeWise)))
	mux.Handle("/getdata", middleware.AuthenticateData(http.HandlerFunc(user.GetData)))
	return mux
}
