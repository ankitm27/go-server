package usecase

import (
	"context"
	"encoding/json"
	"fmt"
	"go-server/database"
	"go-server/elasticSearch"
	"go-server/middleware"
	"net/http"
)

func GetDataTypeWise(w http.ResponseWriter, r *http.Request) {
	data := r.URL.Query()
	fmt.Println("data", data)
	userId := data["userId"][0]
	typeDataStr := ""
	_, ok := data["type"]
	// fmt.Println("value", value)
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
	_, ok = data["userId"]
	// fmt.Println("value", value)
	if ok && data["userId"][0] != "" {
		query["userid"] = userId
	}
	typeData := database.GetData(query, project)
	// fmt.Println("data", typeData)
	message, _ := json.Marshal(typeData)
	// message := "check"
	// w.Write([]byte(message))
	middleware.SendResponseObject(message, w)
}

func GetData(w http.ResponseWriter, r *http.Request) {
	data := r.URL.Query()
	// fmt.Println("data", data)
	// typeData := data["type"][0]
	query := make(map[string]string)
	_, ok := data["type"]
	// fmt.Println("value1111", value)
	// typeData := data["type"][0]
	if ok && data["type"][0] != "" {
		query["datatype"] = data["type"][0]
	}
	_, ok = data["ip"]
	if ok && data["ip"][0] != "" {
		query["ip"] = data["ip"][0]
	}
	_, ok = data["reqId"]
	if ok && data["reqId"][0] != "" {
		query["reqid"] = data["reqId"][0]
	}
	// fmt.Println("query1111", query)
	// fmt.Println("value", value)
	typeDataResult := database.GetUserData(query)
	// fmt.Println("type data result", typeDataResult)
	// w.Write([]byte(typeDataResult))
	message, _ := json.Marshal(typeDataResult)
	// w.Write([]byte(message))
	middleware.SendResponseObject(message, w)
}

func CreateIndex(w http.ResponseWriter, r *http.Request) {
	// fmt.Println("check")
	ctx := context.Background()
	elasticSearch.Ping(ctx, "https://29wd348tb4:ykn6pei12j@test-2867785266.ap-southeast-2.bonsaisearch.net:443")
}

func GetSearchData(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	data := elasticSearch.GetAll(ctx)
	// fmt.Println("data111", data)
	// w.Write([]byte("check"))
	dataObj, err := json.Marshal(data)
	if err != nil {
		fmt.Println("There is some problem, Please try after some time", err)
	}
	// w.Write([]byte(dataObj))
	middleware.SendResponseObject(dataObj, w)
}

func InsertData(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	// data := map[string]string{
	// 	"check":"check",
	// }
	// type data struct {
	// 	check string
	// }
	// var dataObj data
	// dataObj.check = "check"
	elasticSearch.CreateIndexIfDoesNotExist(ctx, "users_index8")
	elasticSearch.InsertUsers(ctx, "users_index8", "user8")
	// fmt.Println("data", data1)
	// w.Write([]byte("check"))
	middleware.SendResponseMessage("data added successfully", w)
}

func GetDemandData(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	_, ok := query["APIURL"]
	queryData := make(map[string]string)
	// fmt.Println("query", query)
	if !ok {
		fmt.Println("There is some problem, Please try after some time1")
		// w.Header().Set("Content-type", "application/json")
		// responseObject := map[string]string{
		// 	"msg": "Please provide all the valid fields",
		// }
		// responseMessage, err := json.Marshal(responseObject)
		// if err != nil {
		// 	fmt.Println("There is some problem, Please try after some time")
		// }
		// w.Write(responseMessage)
		middleware.SendResponseMessage("Please provide all the valid fields", w)
	} else {
		queryData["APIURL"] = query["APIURL"][0]
		ctx := context.Background()
		data := elasticSearch.GetAllSearchData(ctx, queryData)
		// fmt.Println("data", data)
		// fmt.Println("value", value)
		// dataObj, _ := json.Marshal(data)
		// if err != nil {
		// 	fmt.Println("There is some problem, Please try after some time", err)
		// }
		// w.Write([]byte(dataObj))
		middleware.SendResponseObject(data, w)
	}
}
