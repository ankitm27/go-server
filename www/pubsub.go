package www

import redisClient "go-server/redis"

func RunServer() {
	redisClient.RedisClient()
}
