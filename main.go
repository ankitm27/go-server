package main

import (
	database "go-server/database"
	redisclient "go-server/redis"
	socketserver "go-server/socket"
)

func main() {
	redisclient.RedisClient("localhost")
	database.DatabaseConnect()
	// redisclient.Scheduler(1*time.Second)
	socketserver.CreateServer(3333)
}
