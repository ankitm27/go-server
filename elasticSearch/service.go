package elasticSearch

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"

	elasticapi "gopkg.in/olivere/elastic.v5"
)

func CreateIndexIfDoesNotExist(ctx context.Context, indexName string) error {
	client, err := NewElasticClient(context.Background(), false, -1)
	if err != nil {
		fmt.Println("There is some problem, Please try after some time")
	}
	fmt.Println("check1111111")
	exists, err := client.IndexExists(indexName).Do(ctx)
	if err != nil {
		fmt.Println("Thers is some problem, Please try after some time", err)
		return err
	}

	fmt.Println("exists", exists)
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

func Ping(ctx context.Context, client *elasticapi.Client, url string) error {

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

func InsertUsers(ctx context.Context, indexName string, docType string, data interface{}) {
	// insert data in elasticsearch
	client, err := NewElasticClient(context.Background(), false, -1)
	if err != nil {
		fmt.Println("There is some problem, Please try after some time")
	}
	// type User struct {
	// 	UserID    int    `json:"id"`
	// 	Email     string `json:"email"`
	// 	FirstName string `json:"firstname"`
	// 	LastName  string `json:"lastname"`
	// 	Age       int    `json:"age"`
	// 	IsActive  bool   `json:"isActive"`
	// 	Balance   int    `json:"balance"`
	// 	Phone     string `json:"phone"`
	// }
	// var listUsers []User
	// for index := 1; index < 5; index++ {

	// 	user := User{
	// 		UserID:    index,
	// 		Email:     fmt.Sprintf("test%d@gmail.com", index),
	// 		FirstName: fmt.Sprintf("FirstName_%d", index),
	// 		LastName:  fmt.Sprintf("LastName_%d", index),
	// 	}

	// 	listUsers = append(listUsers, user)
	// }

	// for _, userObj := range listUsers {
	_, err = client.Index().Index(indexName).Type(docType).BodyJson(data).Do(ctx)
	if err != nil {
		fmt.Printf("UserId=%d was nos created. Error : %s \n", err)
		// continue
	}
	// }

	// Flush data (need for refreshing data in index) after this command possible to do get.
	client.Flush().Index(indexName).Do(ctx)
}

func GetAll(ctx context.Context, indexName string) interface{} {
	client, err := NewElasticClient(context.Background(), false, -1)
	if err != nil {
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

	fmt.Println("search result", searchResult)

	return convertSearchResultToUsers(searchResult)

}

type User struct {
	UserID    int    `json:"id"`
	Email     string `json:"email"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Age       int    `json:"age"`
	IsActive  bool   `json:"isActive"`
	Balance   int    `json:"balance"`
	Phone     string `json:"phone"`
}

func convertSearchResultToUsers(searchResult *elasticapi.SearchResult) interface{} {
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
	if result != nil {
		return result
	}

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
