package redis

import (
	"fmt"
	"go-server/database"
	"reflect"
	"strings"
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
	err := client1.RPush(key, value).Err()
	if err != nil {
		fmt.Println("error in setting the values", err)
	}
}

func getRedisData(key string) {
	var count int64 = 2
	if getKeyLength(key) >= int(count) {
		data := getNElement(key, count)
		if createEntries(data) {
			removeNElement(key, count)
		}
		// fmt.Println("data", data)
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
	fmt.Println("data", data)
	// createRedisQueue("check", data)
	getRedisData("check1")
	// database.GetDataFromCollection()
}

func createEntries(entries []string) bool {
	fmt.Println("entries", entries)
	dataSlice := convertToSliceObject(entries)
	fmt.Println("data slice", dataSlice)
	database.InsertIntoDb(dataSlice)
	return false
}

func convertToSliceObject(data []string) []interface{} {
	slices := make([]interface{}, len(data))
	j := 0
	for i := 0; i < len(data); i++ {
		// fmt.Println("data i", data[i])
		if isTransportOver(data[i]) {
			// slices[j] = interface{}(data[i])
			// j++
			data[i] = strings.Replace(data[i], "\r\n\r\n", "", 1)
		}
		// fmt.Println("data", data[i])
		slices[j] = interface{}(data[i])
		j++
	}
	return slices
}

func Schedule(interval time.Duration) *time.Ticker {
	ticker := time.NewTicker(interval)
	go func() {
		for range ticker.C {
			getRedisData("check1")
		}
	}()
	return ticker
}

func isTransportOver(data string) (over bool) {
	over = strings.HasSuffix(data, "\r\n\r\n")
	return over
}
