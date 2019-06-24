package services

import (
	"fmt"
	redisClient "go-server/redis"
	"time"
)

func Schedule(interval time.Duration) *time.Ticker {
	fmt.Println("check")
	ticker := time.NewTicker(interval)
	go func() {
		fmt.Println("check1111")
		for range ticker.C {
			fmt.Println("check1111111111111")
			redisClient.ScheduleFunc()
		}
	}()
	return ticker
}
