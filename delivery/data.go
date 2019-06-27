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
	mux.Handle("/createindex", http.HandlerFunc(useCase.CreateIndex))
	mux.Handle("/getsearchdata", middleware.AuthenticateData(http.HandlerFunc(useCase.GetSearchData)))
    mux.Handle("/insertdata",middleware.AuthenticateData(http.HandlerFunc(useCase.InsertData)))
	mux.Handle("/getdemanddata",middleware.AuthenticateData(http.HandlerFunc(useCase.GetDemandData)))
}
