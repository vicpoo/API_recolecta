package core

import (
	"context"
	"fmt"
	"os"
	"sync"
	"time"

	goredis "github.com/redis/go-redis/v9"
)

var (
	redisClient *goredis.Client
	redisOnce   sync.Once
	redisErr    error
)

func ConnectRedis() (*goredis.Client, error) {
	redisOnce.Do(func() {
		addr := fmt.Sprintf("%s:%s", os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT"))
		client := goredis.NewClient(&goredis.Options{
			Addr:         addr,
			Password:     os.Getenv("REDIS_PASSWORD"),
			DialTimeout:  5 * time.Second,
			ReadTimeout:  3 * time.Second,
			WriteTimeout: 3 * time.Second,
		})

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if _, err := client.Ping(ctx).Result(); err != nil {
			redisErr = fmt.Errorf("error connecting to Redis at %s: %w", addr, err)
			return
		}

		redisClient = client
	})
	return redisClient, redisErr
}

func GetRedis() *goredis.Client {
	client, err := ConnectRedis()
	if err != nil {
		panic(fmt.Sprintf("Error al conectar a Redis: %v", err))
	}
	return client
}

func CloseRedis() error {
	if redisClient != nil {
		return redisClient.Close()
	}
	return nil
}
