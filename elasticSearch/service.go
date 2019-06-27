package elasticSearch

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"

	elasticapi "gopkg.in/olivere/elastic.v5"
)

type data struct {
	check string
}

const (
	indexName = "users_index6"
	docType   = "user6"
)

func CreateIndexIfDoesNotExist(ctx context.Context, indexName string) error {
	client, err := NewElasticClient(context.Background(), false, -1)
	if err != nil {
		fmt.Println("There is some problem, Please try after some time")
	}
	// fmt.Println("check1111111")
	exists, err := client.IndexExists(indexName).Do(ctx)
	if err != nil {
		fmt.Println("Thers is some problem, Please try after some time", err)
		return err
	}

	// fmt.Println("exists", exists)
	if exists {
		return nil
	}

	res, err := client.CreateIndex(indexName).Do(ctx)

	if err != nil {
		fmt.Println("There is some problem, Please try after some time", err)
		return err
	}

	if !res.Acknowledged {
		return errors.New("CreateIndex was not acknowledged. Check that timeout value is correct.")
	}
	fmt.Println("res", res)

	return nil
}

func Ping(ctx context.Context, url string) error {
	// Ping the Elasticsearch server to get HttpStatus, version number
	client, err := NewElasticClient(context.Background(), false, -1)
	if err != nil {
		fmt.Println("There is some problem, Please try after some time")
	}
	if client != nil {
		info, code, err := client.Ping(url).Do(ctx)
		if err != nil {
			return err
		}

		fmt.Printf("Elasticsearch returned with code %d and version %s \n", code, info.Version.Number)
		return nil
	}
	return errors.New("elastic client is nil")
}

// func InsertUsers(ctx context.Context, indexName string, docType string, data1 string) {
// 	// insert data in elasticsearch
// 	client, err := NewElasticClient(context.Background(), false, -1)
// 	if err != nil {
// 		fmt.Println("There is some problem, Please try after some time", err)
// 	}
// 	// type User struct {
// 	// 	UserID    int    `json:"id"`
// 	// 	Email     string `json:"email"`
// 	// 	FirstName string `json:"firstname"`
// 	// 	LastName  string `json:"lastname"`
// 	// 	Age       int    `json:"age"`
// 	// 	IsActive  bool   `json:"isActive"`
// 	// 	Balance   int    `json:"balance"`
// 	// 	Phone     string `json:"phone"`
// 	// }
// 	// var listUsers []User
// 	// for index := 1; index < 5; index++ {

// 	// 	user := User{
// 	// 		UserID:    index,
// 	// 		Email:     fmt.Sprintf("test%d@gmail.com", index),
// 	// 		FirstName: fmt.Sprintf("FirstName_%d", index),
// 	// 		LastName:  fmt.Sprintf("LastName_%d", index),
// 	// 	}

// 	// 	listUsers = append(listUsers, user)
// 	// }

// 	// for _, userObj := range listUsers {
// 	dataObj := &data{data1}
// 	fmt.Println("data obj", dataObj)
// 	result, err := client.Index().Index(indexName).Type(docType).BodyJson(dataObj).Do(ctx)
// 	if err != nil {
// 		fmt.Printf("UserId=%d was nos created. Error : %s \n", err)
// 		// continue
// 	}
// 	// }

// 	// Flush data (need for refreshing data in index) after this command possible to do get.
// 	fmt.Println("result", result)
// 	fmt.Println("reflect typeof", reflect.TypeOf(result))
// 	// client.Flush().Index(indexName).Do(ctx)

// }

// func GetAll(ctx context.Context, indexName string) []data {
// 	client, err := NewElasticClient(context.Background(), false, -1)
// 	if err != nil {
// 		fmt.Println("There is some problem, Please try after some time")
// 	}
// 	query := elasticapi.MatchAllQuery{}

// 	searchResult, err := client.Search().
// 		Index(indexName). // search in index
// 		Query(query).     // specify the query
// 		Do(ctx)           // execute
// 	if err != nil {
// 		fmt.Printf("Error during execution GetAll : %s", err.Error())
// 	}

// 	fmt.Println("search result", searchResult)

// 	return convertSearchResultToUsers(searchResult)

// }

// type User struct {
// 	// UserID    int    `json:"id"`
// 	// Email     string `json:"email"`
// 	// FirstName string `json:"firstname"`
// 	// LastName  string `json:"lastname"`
// 	// Age       int    `json:"age"`
// 	// IsActive  bool   `json:"isActive"`
// 	// Balance   int    `json:"balance"`
// 	// Phone     string `json:"phone"`
// 	data         string  `json:"data"` 
// }

// func convertSearchResultToUsers(searchResult *elasticapi.SearchResult) []data {
// 	var result []data
// 	for _, hit := range searchResult.Hits.Hits {
// 		var userObj data
// 		err := json.Unmarshal(*hit.Source, &userObj)
// 		if err != nil {
// 			log.Printf("Can't deserialize 'user' object : %s", err.Error())
// 			continue
// 		}
// 		fmt.Println("user obj", userObj)
// 		result = append(result, userObj)
// 	}
// 	fmt.Println("result", result)
// 	return result
// }

func GetUserByID(ctx context.Context, userID int, indexName string) interface{} {
	client, err := NewElasticClient(context.Background(), false, -1)
	if err != nil {
		fmt.Println("There is some problem, Please try after some time")
	}
	query := elasticapi.NewBoolQuery()
	//sortObj := elasticapi.NewFieldSort("creation_date").Desc()
	musts := []elasticapi.Query{elasticapi.NewTermQuery("id", userID)}
	query = query.Must(musts...)

	searchResult, err := client.Search().
		Index(indexName). // search in index
		Query(query).     // specify the query
		Do(ctx)           // execute
	if err != nil {
		fmt.Printf("Error during execution FindAndPrintUsers : %s", err.Error())
	}

	var result = convertSearchResultToUsers(searchResult)
	// if result != nil {
	// 	return result
	// }
	fmt.Println("result", result)
	return User{}
}

func GetAllActiveUsers(ctx context.Context, indexName string) interface{} {
	client, err := NewElasticClient(context.Background(), false, -1)
	if err != nil {
		fmt.Println("There is some problem, Please try after some time")
	}
	query := elasticapi.NewBoolQuery()
	query = query.Must(elasticapi.NewTermQuery("isActive", true))

	searchResult, err := client.Search().
		Index(indexName). // search in index
		Query(query).     // specify the query
		Do(ctx)           // execute
	if err != nil {
		fmt.Printf("Error during execution GetAll : %s", err.Error())
	}

	return convertSearchResultToUsers(searchResult)

}

func DeleteUser(ctx context.Context, userID int, indexName string, docType string) {
	client, err = NewElasticClient(context.Background(), false, -1)
	if err != nil {
		fmt.Println("There is some problem, Please try after some time")
	}
	bq := elasticapi.NewBoolQuery()
	bq.Must(elasticapi.NewTermQuery("id", userID))

	_, err := elasticapi.NewDeleteByQueryService(client).Index(indexName).Type(docType).Query(bq).Do(ctx)
	if err != nil {
		fmt.Printf("Error during execution DeleteUser : %s", err.Error())
		return
	}

	// Flush data (need for refreshing data in index) after this command possible to do get.
	client.Flush().Index(indexName).Do(ctx)
}

// func InsertUsers(ctx context.Context, indexName string, docType string) {
// 	// insert data in elasticsearch
// 	client, err := NewElasticClient(context.Background(), false, -1)
// 	if err != nil {
// 		fmt.Println("There is some problem, Please try after some time")
// 	}
// 	// var listUsers []User
// 	// for index := 1; index < 5; index++ {

// 	user := User{
// 		// UserID:    22,
// 		// Email:     fmt.Sprintf("test1212%d@gmail.com", 1),
// 		// FirstName: fmt.Sprintf("FirstName_%d", 1),
// 		// LastName:  fmt.Sprintf("LastName_%d", 1),
// 		data:          "check",
// 	}

// 	// listUsers = append(listUsers, user)
// 	// }

// 	// for _, userObj := range listUsers {
// 	fmt.Println("data1111", user)
// 	_, err = client.Index().Index(indexName).Type(docType).BodyJson(user).Do(ctx)
// 	if err != nil {
// 		fmt.Printf("UserId=%d was nos created. Error : %s \n", err.Error())
// 		// continue
// 	}
// 	// }

// 	// Flush data (need for refreshing data in index) after this command possible to do get.
// 	client.Flush().Index(indexName).Do(ctx)
// }

// func GetAll(ctx context.Context) []User {
// 	client, err := NewElasticClient(context.Background(), false, -1)
// 	if err != nil {
// 		fmt.Println("There is some problem, Please try after some time")
// 	}
// 	query := elasticapi.MatchAllQuery{}

// 	searchResult, err := client.Search().
// 		Index(indexName). // search in index
// 		Query(query).     // specify the query
// 		Do(ctx)           // execute
// 	if err != nil {
// 		fmt.Printf("Error during execution GetAll : %s", err.Error())
// 	}

// 	return convertSearchResultToUsers(searchResult)
// }

// func convertSearchResultToUsers(searchResult *elasticapi.SearchResult) []User {
// 	var result []User
// 	for _, hit := range searchResult.Hits.Hits {
// 		var userObj User
// 		err := json.Unmarshal(*hit.Source, &userObj)
// 		if err != nil {
// 			log.Printf("Can't deserialize 'user' object : %s", err.Error())
// 			continue
// 		}
// 		result = append(result, userObj)
// 	}
// 	return result
// }

type User struct {
	// UserID    int    `json:"id"`
	// Email     string `json:"email"`
	// FirstName string `json:"firstname"`
	// LastName  string `json:"lastname"`
	// Age       int    `json:"age"`
	// IsActive  bool   `json:"isActive"`
	// Balance   int    `json:"balance"`
	// Phone     string `json:"phone"`
    Data      string `json:data`
}

func InsertUsers(ctx context.Context,indexName string,docType string) {
	// insert data in elasticsearch
	client,err := NewElasticClient(context.Background(),false,-1)
	if(err != nil){
		fmt.Println("There is some problem, Please try after some time",err)
	}
	// var listUsers []User
	// for index := 1; index < 5; index++ {

		user := User{
			// UserID:    122123,
			// Email:     fmt.Sprintf("test%d@gmail.com", 1222112),
			// FirstName: fmt.Sprintf("FirstName_%d", 12121231),
			// LastName:  fmt.Sprintf("LastName_%d", 112112),
		    Data:       "check",
		}

		// listUsers = append(listUsers, user)
	// }
       fmt.Println("user",user) 
	// for _, userObj := range listUsers {
		_, err = client.Index().Index(indexName).Type(docType).BodyJson(user).Do(ctx)
		if err != nil {
			fmt.Printf("UserId=%d was nos created. Error : %s \n",  err.Error())
			// continue
		}
	// }

	// Flush data (need for refreshing data in index) after this command possible to do get.
	client.Flush().Index(indexName).Do(ctx)
}

func GetAll(ctx context.Context) []User {
	client,err := NewElasticClient(context.Background(),false,-1)
	if(err != nil){
		fmt.Println("There is some problem, Please try after some time")
	}
	query := elasticapi.MatchAllQuery{}

	searchResult, err := client.Search().
		Index(indexName). // search in index
		Query(query).     // specify the query
		Do(ctx)           // execute
	if err != nil {
		fmt.Printf("Error during execution GetAll : %s", err.Error())
	}

	return convertSearchResultToUsers(searchResult)
}

func convertSearchResultToUsers(searchResult *elasticapi.SearchResult) []User {
	var result []User
	for _, hit := range searchResult.Hits.Hits {
		var userObj User
		err := json.Unmarshal(*hit.Source, &userObj)
		if err != nil {
			log.Printf("Can't deserialize 'user' object : %s", err.Error())
			continue
		}
		result = append(result, userObj)
	}
	return result
}
