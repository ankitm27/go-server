package main

import (
	database "go-server/database"
	redisclient "go-server/redis"
	socketserver "go-server/socket"
)

func main() {
	redisclient.RedisClient("localhost")
	database.DatabaseConnect()
	socketserver.CreateServer(3333)

}
