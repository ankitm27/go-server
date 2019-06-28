package middleware

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func SendResponseMessage(message string, w http.ResponseWriter) (int, error) {
	// responseObject := map[string]string{
	// 	"msg": message,
	//     "status":true,
	// }
	responseObject := make(map[string]interface{})
	responseObject["msg"] = message
	responseObject["status"] = true
	responseObject["statusCode"] = 200
	responseMessage, err := json.Marshal(responseObject)
	if err != nil {
		fmt.Println("There is some problem, Please try after some time", err)
	}
	w.Header().Set("content-type", "json/application")
	return w.Write(responseMessage)
}

func SendResponseObject(object interface{}, w http.ResponseWriter) (int, error) {
	// responseMessage, err := json.Marshal(object)
	// if err != nil {
	// 	fmt.Println("There is some problem, Please try after some time")
	// }
	res := map[string]interface{}{
		"status":     true,
		"statusCode": 200,
		"data":       object,
	}
	responseMessage, err := json.Marshal(res)
	if err != nil {
		fmt.Println("There is some problem, Please try after some time")
	}
	w.Header().Set("content-type", "json/application")
	return w.Write([]byte(responseMessage))
}
