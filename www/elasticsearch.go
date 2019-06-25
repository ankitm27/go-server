package www

import (
	"context"
	"fmt"
	elasticSearch "go-server/elasticSearch"
)

func StartElasticSearch() {
	// ctx := context.Background()
	fmt.Println("check")
	elasticSearch, err := elasticSearch.NewElasticClient(context.Background(),false, -1)
	if err != nil {
		fmt.Println("There is some problem, Please try after some time", err)
	}
	fmt.Println("elastic search", elasticSearch)
}
