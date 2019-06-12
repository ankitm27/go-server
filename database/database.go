package database

import (
	"context"
	"fmt"
	"log"
	"net/url"
	"reflect"
	"regexp"
	"time"

	cryptography "go-server/cryptography"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

var client *mongo.Client
var databaseConnection = false

type User struct {
	ID       string `bson:"_id" json:"_id,omitempty"`
	Email    string `bson:"email" json:"email,omitempty"`
	Password string `bson:"password" json:"password,omitempty"`
	Key      string `bson:"key" json:"key,omitempty"`
	Secret   string `bson:"secret" json:"secret,omitempty"`
}

func (user User) Validate() (errs url.Values) {
	regexpEmail := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	if user.Email == "" {
		errs.Add("email", "This is required field")
	}
	regexpEmail.MatchString(user.Email)
	if !regexpEmail.MatchString(user.Email) {
		errs.Add("email", "The email field should be a valid email address!")
	}
	if user.Password == "" {
		errs.Add("password", "This is required field")
	}
	return errs
}

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

func InsertIntoDb(data []interface{}) []interface{} {
	client1 := DatabaseConnect()
	collection := client1.Database("testing").Collection("numbers")
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	// datum = make(interface{}, map({"Name":"Eve","Age":6,"Parents":"Alice"})

	// _, err := collection.InsertOne(ctx, bson.M{"Name": "Eve", "Age": 6, "Parents": "Alice"})
	results, err := collection.InsertMany(ctx, data)

	if err != nil {
		fmt.Println("error in saving the data", err)
		return nil
	}
	fmt.Println(results.InsertedIDs)
	return results.InsertedIDs
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

func GetUser(data map[string]string) *User {
	result := &User{}
	client = DatabaseConnect()
	collection := client.Database("testing").Collection("users")
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	projection := map[string]int{
		"_id":      0,
		"password": 0,
	}
	err := collection.FindOne(ctx, interface{}(data), options.FindOne().SetProjection(interface{}(projection))).Decode(result)

	if err != nil {
		fmt.Println("Error", err)
	}
	return result
}

func CreateUser(data map[string]string) (interface{}, error) {
	client = DatabaseConnect()
	fmt.Println(data)
	key, keyerr := cryptography.Encrypt(data["email"])
	if keyerr != nil {
		panic(keyerr)
	}
	secret, secreterr := cryptography.Encrypt(data["password"])
	if secreterr != nil {
		panic(secreterr)
	}
	userData := User{
		ID:       bson.NewObjectId().Hex(),
		Email:    data["email"],
		Password: data["password"],
		Key:      key,
		Secret:   secret,
	}
	fmt.Println(userData.Validate())

	collection := client.Database("testing").Collection("users")
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	result, err := collection.InsertOne(ctx, userData)
	if err != nil {
		return "", err
	}
	return result.InsertedID, nil
}
