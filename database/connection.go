package database

import (
	"context"
	"fmt"
	"time"

	utility "go-server/utility"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client1 *mongo.Client
var err error

func DatabaseConnect() *mongo.Client {
	// fmt.Println("check")
	if client1 != nil {
		return client1
	} else {
		// fmt.Println("check1")
		ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
		config := utility.GetConfig()
		// fmt.Println("client", client1)
		client1, err = mongo.Connect(ctx, options.Client().ApplyURI(config.DatabaseUrl))
		if err != nil {
			fmt.Println("error in connecting with mongo", err)
		}
		// fmt.Println(reflect.TypeOf(client))
		databaseConnection = true
		// }
		return client1
	}
}
