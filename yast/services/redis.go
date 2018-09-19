package services

import "github.com/go-redis/redis"

// RedisClient Exported redis client
var RedisClient *redis.Client

func init() {
	connectionString := ApplicationConfig.Redis.Host
	client := redis.NewClient(&redis.Options{
		Addr:     connectionString,
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	pong, err := client.Ping().Result()
	if err != nil {
		Logger.Error("Error connecting redis", err)
	} else {
		Logger.Info("Connection successful", pong)
	}
	RedisClient = client
}
