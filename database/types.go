package database

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetData(userId map[string]string, project map[string]int) *DataType {
	// return "check"
	// data := make(map[string]int)
	// data["info"] = 1
	// data["success"] = 5
	// data["warning"] = 2
	// data["error"] = 1
	// return data
	client = DatabaseConnect()
	collection := client.Database("testing").Collection("typedata")
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	result := &DataType{}
	// fmt.Println("user id", userId)
	// data := DataType{
	// 	UserId:  "check",
	// 	Success: "1",
	// 	Info:    "1",
	// 	Warning: "1",
	// 	Error:   "1",
	// }
	// result1, err := collection.InsertOne(ctx, data)
	// fmt.Println("result", result1)
	err := collection.FindOne(ctx, interface{}(userId), options.FindOne().SetProjection(interface{}(project))).Decode(result)
	fmt.Println("result", result)
	if err != nil {
		fmt.Println("There is some problem in fetching the data", err)
	}
	return result
}
