package delivery

import "github.com/gorilla/mux"

func Delivery() *mux.Router {
	mux := mux.NewRouter()
	UserDelivery(mux)
	DataDelivery(mux)
	return mux
}
