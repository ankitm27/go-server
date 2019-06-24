package delivery

import (
	"go-server/middleware"
	"net/http"

	useCase "go-server/useCase"

	"github.com/gorilla/mux"
)

// func DataDelivery(mux) *mux.Router {
// 	mux.Handle("/getdatatypewise", middleware.AuthenticateData(http.HandleFunc(useCase.GetDataTypeWise)))
// }

func DataDelivery(mux *mux.Router) {
	mux.Handle("/getdatatypewise", middleware.AuthenticateData(http.HandlerFunc(useCase.GetDataTypeWise)))
	mux.Handle("/getdata", middleware.AuthenticateData(http.HandlerFunc(useCase.GetData)))
}
