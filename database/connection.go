package database

import (
	"context"
	"fmt"
	"time"

	utility "go-server/utility"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func DatabaseConnect() *mongo.Client {
	// fmt.Println("check")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	config := utility.GetConfig()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(config.DatabaseUrl))
	if err != nil {
		fmt.Println("error in connecting with mongo", err)
	}
	// fmt.Println(reflect.TypeOf(client))
	databaseConnection = true
	// }
	return client
}
