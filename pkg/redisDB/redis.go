package redisDB

import (
	"context"
	"github.com/redis/go-redis/v9"
	"log"
	"os"
)

func NewRedisDB() (*redis.Client, error) {
	opt, err := redis.ParseURL(os.Getenv("REDIS_URL"))
	if err != nil {
		panic(err)
	}

	client := redis.NewClient(opt)
	_, err = client.Ping(context.Background()).Result()
	if err != nil {
		log.Println("Error connecting to Redis:", err)
		return nil, err
	}
	log.Println("Connected to Redis")
	return client, nil
}
