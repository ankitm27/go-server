package database

import (
	"context"
	"fmt"
	"time"
)

func GetUserData(typeData map[string]string) *Data {
	// return "check"
	client = DatabaseConnect()
	collection := client.Database("testing").Collection("data")
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	result := &Data{}
	err := collection.FindOne(ctx, interface{}(typeData)).Decode(result)
	if err != nil {
		fmt.Println("There is some problem in fetching the data", err)
		var data *Data
		return data
	}
	fmt.Println("result", result)
	return result
}
