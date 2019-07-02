package database

import (
	"context"
	"fmt"
	"net/url"
	"regexp"
	"time"

	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

func GetUser(data map[string]string) *User {
	result := &User{}
	client = DatabaseConnect()
	collection := client.Database("testing").Collection("users")
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	projection := map[string]int{
		"_id":            0,
		"password":       0,
		"hashedPassword": 0,
	}
	err := collection.FindOne(ctx, interface{}(data), options.FindOne().SetProjection(interface{}(projection))).Decode(result)

	if err != nil {
		fmt.Println("Error 1212", err)
	}
	return result
}

func CreateUser(data map[string]string) (interface{}, error) {
	client = DatabaseConnect()
	// fmt.Println(data)
	// key, keyerr := cryptography.Encrypt(data["email"])
	// if keyerr != nil {
	// 	// panic(keyerr)
	// 	fmt.Println("key err", keyerr)
	// }
	// secret, secreterr := cryptography.Encrypt(data["password"])
	// if secreterr != nil {
	// 	// panic(secreterr)
	// 	fmt.Println("secret err", secreterr)
	// }
	userData := User{
		ID:             bson.NewObjectId().Hex(),
		Email:          data["email"],
		Password:       data["password"],
		Key:            data["key"],
		Secret:         data["secret"],
		HashedPassword: data["hashedPassword"],
	}
	// fmt.Println(userData.Validate())
	fmt.Println("user data", userData)
	userData.validate()
	collection := client.Database("testing").Collection("users")
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	result, err := collection.InsertOne(ctx, userData)
	if err != nil {
		return "", err
	}
	// fmt.Println("result", result)
	return result.InsertedID, nil
}

func IsSecretValid(authenticate map[string]string) bool {
	client = DatabaseConnect()
	collection := client.Database("testing").Collection("users")
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	result := &User{}
	// fmt.Println("autheticate", authenticate)
	err := collection.FindOne(ctx, interface{}(authenticate)).Decode(result)
	if err != nil {
		fmt.Println("There is some problem in finding the secret, Please try after some time", err)
		return false
	}
	return true
}

func (user User) validate() (errs url.Values) {
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
