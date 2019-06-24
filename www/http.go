package www

import (
	"fmt"
	"go-server/delivery"
	utility "go-server/utility"

	"github.com/fvbock/endless"
)

var config1 = utility.GetConfig()
var backendUrl = config.BackendUrl + ":" + config1.Port

func HttpServer() {
	mux := delivery.Delivery()
	fmt.Println("backend url", backendUrl)
	err := endless.ListenAndServe(backendUrl, mux)
	if err != nil {
		fmt.Println("there are some error in starting the golang server", err)
	}
}
