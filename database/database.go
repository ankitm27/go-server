package database

import (
	"context"
	"fmt"
	"log"
	"reflect"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

var client *mongo.Client
var databaseConnection = false

func DatabaseConnect() *mongo.Client {
	fmt.Println("check")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		fmt.Println("error in connecting with mongo", err)
	}
	fmt.Println(reflect.TypeOf(client))
	databaseConnection = true
	// }
	return client
}

func InsertIntoDb(data []string) {
	// fmt.Println("client", client)
	client1 := DatabaseConnect()
	collection := client1.Database("testing").Collection("numbers")
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	// fmt.Println(reflect.TypeOf(bson.M(data)))
	res, err := collection.InsertOne(ctx, bson.M{"name": "pi", "value": 3.14159})
	// collection.InsertMany(ctx, bson.Mdata)
	if err != nil {
		fmt.Println("error in saving the data", err)
	}
	id := res.InsertedID
	fmt.Println("id", id)
	// fmt.Println("ctx", ctx)
	// fmt.Println("res", res)
}

func GetDataFromCollection() {
	var result struct {
		Value float64
	}
	client1 := DatabaseConnect()
	collection := client1.Database("testing").Collection("numbers")

	filter := bson.M{"name": "pi"}
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	err := collection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("result", result)
}
