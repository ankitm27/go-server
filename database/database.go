package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"go-server/Models"

	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
)

var client *mongo.Client
var databaseConnection = false

// type User struct {
// 	ID       string `bson:"_id" json:"_id,omitempty"`
// 	Email    string `bson:"email" json:"email"`
// 	Password string `bson:"password" json:"-"`
// 	Key      string `bson:"key" json:"key"`
// 	Secret   string `bson:"secret" json:"secret"`
// }

type User Models.User
type DataType Models.TypeData
type Data Models.Data

// func (user User) Validate() (errs url.Values) {
// 	regexpEmail := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
// 	if user.Email == "" {
// 		errs.Add("email", "This is required field")
// 	}
// 	regexpEmail.MatchString(user.Email)
// 	if !regexpEmail.MatchString(user.Email) {
// 		errs.Add("email", "The email field should be a valid email address!")
// 	}
// 	if user.Password == "" {
// 		errs.Add("password", "This is required field")
// 	}
// 	return errs
// }

// func DatabaseConnect() *mongo.Client {
// 	// fmt.Println("check")
// 	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
// 	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
// 	if err != nil {
// 		fmt.Println("error in connecting with mongo", err)
// 	}
// 	// fmt.Println(reflect.TypeOf(client))
// 	databaseConnection = true
// 	// }
// 	return client
// }

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
	// fmt.Println(results.InsertedIDs)
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

// func GetUser(data map[string]string) *User {
// 	result := &User{}
// 	client = DatabaseConnect()
// 	collection := client.Database("testing").Collection("users")
// 	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
// 	projection := map[string]int{
// 		"_id":            0,
// 		"password":       0,
// 		"hashedPassword": 0,
// 	}
// 	err := collection.FindOne(ctx, interface{}(data), options.FindOne().SetProjection(interface{}(projection))).Decode(result)

// 	if err != nil {
// 		fmt.Println("Error 1212", err)
// 	}
// 	return result
// }

// func CreateUser(data map[string]string) (interface{}, error) {
// 	client = DatabaseConnect()
// 	// fmt.Println(data)
// 	// key, keyerr := cryptography.Encrypt(data["email"])
// 	// if keyerr != nil {
// 	// 	// panic(keyerr)
// 	// 	fmt.Println("key err", keyerr)
// 	// }
// 	// secret, secreterr := cryptography.Encrypt(data["password"])
// 	// if secreterr != nil {
// 	// 	// panic(secreterr)
// 	// 	fmt.Println("secret err", secreterr)
// 	// }
// 	userData := User{
// 		ID:             bson.NewObjectId().Hex(),
// 		Email:          data["email"],
// 		Password:       data["password"],
// 		Key:            data["key"],
// 		Secret:         data["secret"],
// 		HashedPassword: data["hashedPassword"],
// 	}
// 	// fmt.Println(userData.Validate())
// 	fmt.Println("user data", userData)
// 	userData.Validate()
// 	collection := client.Database("testing").Collection("users")
// 	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
// 	result, err := collection.InsertOne(ctx, userData)
// 	if err != nil {
// 		return "", err
// 	}
// 	fmt.Println("result", result)
// 	return result.InsertedID, nil
// }

// func GetData(userId map[string]string, project map[string]int) *DataType {
// 	// return "check"
// 	// data := make(map[string]int)
// 	// data["info"] = 1
// 	// data["success"] = 5
// 	// data["warning"] = 2
// 	// data["error"] = 1
// 	// return data
// 	client = DatabaseConnect()
// 	collection := client.Database("testing").Collection("typedata")
// 	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
// 	result := &DataType{}
// 	// fmt.Println("user id", userId)
// 	// data := DataType{
// 	// 	UserId:  "check",
// 	// 	Success: "1",
// 	// 	Info:    "1",
// 	// 	Warning: "1",
// 	// 	Error:   "1",
// 	// }
// 	// result1, err := collection.InsertOne(ctx, data)
// 	// fmt.Println("result", result1)
// 	err := collection.FindOne(ctx, interface{}(userId), options.FindOne().SetProjection(interface{}(project))).Decode(result)
// 	fmt.Println("result", result)
// 	if err != nil {
// 		fmt.Println("There is some problem in fetching the data", err)
// 	}
// 	return result
// }

// func GetUserData(typeData map[string]string) *Data {
// 	// return "check"
// 	client = DatabaseConnect()
// 	collection := client.Database("testing").Collection("data")
// 	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
// 	result := &Data{}
// 	err := collection.FindOne(ctx, interface{}(typeData)).Decode(result)
// 	if err != nil {
// 		fmt.Println("There is some problem in fetching the data", err)
// 		var data *Data
// 		return data
// 	}
// 	fmt.Println("result", result)
// 	return result
// }

// func IsSecretValid(authenticate map[string]string) bool {
// 	client = DatabaseConnect()
// 	collection := client.Database("testing").Collection("users")
// 	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
// 	result := &User{}
// 	// fmt.Println("autheticate", authenticate)
// 	err := collection.FindOne(ctx, interface{}(authenticate)).Decode(result)
// 	if err != nil {
// 		fmt.Println("There is some problem in finding the secret, Please try after some time", err)
// 		return false
// 	}
// 	return true
// }
