package database

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
    utility "go-server/utility"
)

func DatabaseConnect() *mongo.Client {
	// fmt.Println("check")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(utility.DatabaseUrl))
	if err != nil {
		fmt.Println("error in connecting with mongo", err)
	}
	// fmt.Println(reflect.TypeOf(client))
	databaseConnection = true
	// }
	return client
}
