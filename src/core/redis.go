package core

import (
	"context"
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/redis/go-redis/v9"
)

var (
	redisClient *redis.Client
	redisOnce   sync.Once
	redisErr    error
)

func ConnectRedis() (*redis.Client, error) {
	redisOnce.Do(func() {
		host := os.Getenv("REDIS_HOST")
		port := os.Getenv("REDIS_PORT")
		password := os.Getenv("REDIS_PASSWORD")
		if host == "" {
			host = "redis"
		}
		if port == "" {
			port = "6379"
		}

		rdb := redis.NewClient(&redis.Options{
			Addr:         fmt.Sprintf("%s:%s", host, port),
			Password:     password,
			DB:           0,
			DialTimeout:  10 * time.Second,
			ReadTimeout:  5 * time.Second,
			WriteTimeout: 5 * time.Second,
		})

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		if _, err := rdb.Ping(ctx).Result(); err != nil {
			redisErr = fmt.Errorf("error connecting to redis: %w", err)
			return
		}

		redisClient = rdb
	})

	return redisClient, redisErr
}

func GetRedis() *redis.Client {
	rdb, err := ConnectRedis()
	if err != nil {
		panic(fmt.Sprintf("Error al conectar a Redis: %v", err))
	}
	return rdb
}
