package user

import (
	"encoding/json"
	"fmt"
	"go-server/database"
	"io/ioutil"
	"net/http"
)

func SignUp(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println("there is some error in sign up", err)
	}
	payload := make(map[string]string)
	error := json.Unmarshal(data, &payload)
	if error != nil {
		fmt.Println("error", error)
	}
	var message []byte
	user := database.GetUser(map[string]string{"email": payload["email"]})
	if user.ID == "" {
		user, err := database.CreateUser(payload)
		if err != nil {
			fmt.Println("error", err)
			return
		}
		fmt.Println("user1212", user)
		user1 := database.GetUser(map[string]string{"email": payload["email"]})
		fmt.Println("user 1111", user1)
		message, _ = json.Marshal(user1)
	} else if user.Password == payload["Password"] {
		fmt.Println("Invalid password")
	} else {
		message, _ = json.Marshal(user)
	}
	w.Write(message)
}

func GetDataTypeWise(w http.ResponseWriter, r *http.Request) {
	data := r.URL.Query()
	fmt.Println("data", data)
	userId := data["userId"][0]
	typeDataStr := ""
	if data["type"][0] != "" {
		typeDataStr = data["type"][0]
	}
	// payload := make(map[string]string)
	// err := json.Unmarshal(data, &payload)
	// if err != nil {
	// 	fmt.Println("There is some problem in getting data", err)
	// }
	project := make(map[string]int)
	if typeDataStr != "" {
		project[typeDataStr] = 1
	}
	typeData := database.GetData(map[string]string{"userid": userId}, project)
	fmt.Println("data", typeData)
	message, _ := json.Marshal(typeData)
	// message := "check"
	w.Write([]byte(message))
}
