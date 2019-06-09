package main

import (
	redisclient "go-server/redis"
	socketserver "go-server/socket"
)

func main() {
	redisclient.RedisClient("localhost")
	socketserver.CreateServer(3333)

}
