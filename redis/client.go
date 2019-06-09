package redis

import (
	"fmt"
	"go-server/database"
	"reflect"
	"time"

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

func createRedisQueue(key string, value string) {
	client1 := RedisClient("localhost")
	// fmt.Println("client", client)
	err := client1.RPush(key, value).Err()
	if err != nil {
		fmt.Println("error in setting the values", err)
	}
}

func getRedisData(key string) {
	// client1 := RedisClient("localhost")
	var count int64 = 2
	if getKeyLength(key) >= int(count) {
		data := getNElement(key, count)
		if createEntries(data) {
			removeNElement(key, count)
		}
	}
}

func getKeyLength(key string) int {
	client1 := RedisClient("localhost")
	val, err := client1.LLen(key).Result()
	if err != nil {
		fmt.Println("error in getting the data", err)
		return 0
	}
	fmt.Println("val", val)
	return int(val)
}

func getNElement(key string, n int64) []string {
	client1 := RedisClient("localhost")
	val, err := client1.LRange(key, 0, n-1).Result()
	if err != nil {
		fmt.Println("Error while getting elements: ")
		fmt.Println(err)
	}
	return val
}

func removeNElement(key string, n int64) {
	client1 := RedisClient("localhost")
	_, err := client1.LTrim(key, n+1, -1).Result()
	if err != nil {
		fmt.Println("Error while getting elements: ")
		fmt.Println(err)
	}
	// fmt.Println("val", val)
}

func AddDataIntoRedis(data string) {
	// getKeyLength("check1")
	// getRedisData("check1")
	// database.InsertIntoDb()
	createRedisQueue("check", data)
	// database.GetDataFromCollection()
}

// func redisGetter() {
// length := getKeyLength("check1")
// 	if length >= 50 {
// 		getRedisData("check1")
// 		database.InsertIntoDb()
// 	}
// }

func createEntries(entries []string) bool {
	// TODO: Create entries
	// fmt.Println("Entries created for:-")
	// fmt.Println(entries)
	database.InsertIntoDb()
	return true
}

func Schedule(interval time.Duration) *time.Ticker {
	ticker := time.NewTicker(interval)
	go func() {
		for range ticker.C {
			getRedisData("check")
		}
	}()
	return ticker
}
