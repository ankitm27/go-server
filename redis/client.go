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
		client = redis.NewClient(&redis.Options{
			Addr:     host + ":6379",
			Password: "",
			DB:       0,
		})
		fmt.Println(reflect.TypeOf(client))
		isConnectionAvailable = true
	}
	return client

}

func createRedisQuene(key string, value string) {
	client1 := RedisClient("localhost")
	// fmt.Println("client", client)
	err := client1.Set(key, value, 0).Err()
	if err != nil {
		fmt.Println("error in setting the values", err)
	}
	fmt.Println("check")
}

func getRedisData(key string) {
	client1 := RedisClient("localhost")
	val, err := client1.Get(key).Result()
	if err != nil {
		fmt.Println("error in getting the data", err)
	}
	fmt.Println("val", val)

}

func AddDataIntoRedis(data string) {
	// fmt.Println(client, "check", string(data))
	createRedisQuene("check", string(data))
	getRedisData("check")
}
