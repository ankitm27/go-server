package redis

import (
	"fmt"
	"reflect"

	"github.com/go-redis/redis"
)

var client *redis.Client
var isConnectionAvailable = false

func RedisClient(host string) *redis.Client {
	if !isConnectionAvailable {
		client := redis.NewClient(&redis.Options{
			Addr:     host + ":6379",
			Password: "",
			DB:       0,
		})
		fmt.Println(reflect.TypeOf(client))
		isConnectionAvailable = true
	}
	return client
}

func AddDataIntoRedis(data string) {
	fmt.Println(data)
}
