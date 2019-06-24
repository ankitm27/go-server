package usecase

import (
	"encoding/json"
	"fmt"
	"go-server/database"
	"net/http"
)

func GetDataTypeWise(w http.ResponseWriter, r *http.Request) {
	data := r.URL.Query()
	fmt.Println("data", data)
	userId := data["userId"][0]
	typeDataStr := ""
	value, ok := data["type"]
	fmt.Println("value", value)
	if ok && data["type"][0] != "" {
		typeDataStr = data["type"][0]
	}
	// payload := make(map[string]string)
	// err := json.Unmarshal(data, &payload)
	// if err != nil {
	// 	fmt.Println("There is some problem in getting data", err)
	// }
	project := make(map[string]int)
	// fmt.Println("project",)
	if typeDataStr != "" {
		project[typeDataStr] = 1
	}
	query := make(map[string]string)
	value, ok = data["userId"]
	fmt.Println("value", value)
	if ok && data["userId"][0] != "" {
		query["userid"] = userId
	}
	typeData := database.GetData(query, project)
	fmt.Println("data", typeData)
	message, _ := json.Marshal(typeData)
	// message := "check"
	w.Write([]byte(message))
}

func GetData(w http.ResponseWriter, r *http.Request) {
	data := r.URL.Query()
	fmt.Println("data", data)
	// typeData := data["type"][0]
	query := make(map[string]string)
	value, ok := data["type"]
	// fmt.Println("value1111", value)
	// typeData := data["type"][0]
	if ok && data["type"][0] != "" {
		query["datatype"] = data["type"][0]
	}
	value, ok = data["ip"]
	if ok && data["ip"][0] != "" {
		query["ip"] = data["ip"][0]
	}
	value, ok = data["reqId"]
	if ok && data["reqId"][0] != "" {
		query["reqid"] = data["reqId"][0]
	}
	fmt.Println("query1111", query)
	fmt.Println("value", value)
	typeDataResult := database.GetUserData(query)
	fmt.Println("type data result", typeDataResult)
	// w.Write([]byte(typeDataResult))
	message, _ := json.Marshal(typeDataResult)
	w.Write([]byte(message))
}
