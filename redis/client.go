package redis

import (
	"fmt"
	"go-server/database"
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
	err := client1.LPush(key, value).Err()
	if err != nil {
		fmt.Println("error in setting the values", err)
	}
}

func getRedisData(key string) {
	client1 := RedisClient("localhost")
	val, err := client1.LPop(key).Result()
	if err != nil {
		fmt.Println("error in getting the data", err)
	}
	fmt.Println("val", val)
}

func getKeyLength(key string) int64 {
	client1 := RedisClient("localhost")
	val, err := client1.LLen(key).Result()
	if err != nil {
		fmt.Println("error in getting the data", err)
	}
	// fmt.Println("val", val)
	return val
}

func AddDataIntoRedis(data string) {
	createRedisQuene("check1", string(data))
	// getKeyLength("check1")
	// getRedisData("check1")
	database.InsertIntoDb()
	database.GetDataFromCollection()
}

func redisGetter() {
	length := getKeyLength("check1")
	if length >= 50 {
		getRedisData("check1")
		database.InsertIntoDb()
	}
}
