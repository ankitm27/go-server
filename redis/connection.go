package redis

import (
	utility "go-server/utility"

	"github.com/go-redis/redis"
)

var client *redis.Client
var isConnectionAvailable = false

var config = utility.GetConfig()
var url = config.PubSubUrl + ":" + config.PubSubPort

func RedisClient() *redis.Client {
	// fmt.Println("check11111111")
	if !isConnectionAvailable {
		client = redis.NewClient(&redis.Options{
			Addr:     url,
			Password: "",
			DB:       0,
		})
		// fmt.Println(reflect.TypeOf(client))
		// fmt.Println("check1111111111")
		isConnectionAvailable = true
	}
	return client

}
