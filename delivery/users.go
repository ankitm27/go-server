package delivery

import (
	useCase "go-server/useCase"
	"net/http"

	"github.com/gorilla/mux"
)

func UserDelivery(mux *mux.Router) {
	// mux := mux.NewRouter()
	mux.Handle("/signup", http.HandlerFunc(useCase.SignUp))
	// mux.Handle("/getdatatypewise", middleware.AuthenticateData(http.HandlerFunc(useCase.GetDataTypeWise)))
	// mux.Handle("/getdata", middleware.AuthenticateData(http.HandlerFunc(useCase.GetData)))
}
