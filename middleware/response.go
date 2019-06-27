package middleware

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func SendResponseMessage(message string, w http.ResponseWriter) {
	responseObject := map[string]string{
		"msg": message,
	}
	responseMessage, err := json.Marshal(responseObject)
	if err != nil {
		fmt.Println("There is some problem, Please try after some time")
	}
	w.Write(responseMessage)
}

func SendResponseObject(object []byte, w http.ResponseWriter) {
	// responseMessage, err := json.Marshal(object)
	// if err != nil {
	// 	fmt.Println("There is some problem, Please try after some time")
	// }
	w.Write([]byte(object))
}
