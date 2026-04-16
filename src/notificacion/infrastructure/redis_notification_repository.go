package infrastructure

import (
	"context"

	"github.com/redis/go-redis/v9"
)

type RedisNotificationRepository struct {
	rdb *redis.Client
}

func NewRedisNotificationRepository(rdb *redis.Client) *RedisNotificationRepository {
	return &RedisNotificationRepository{rdb: rdb}
}

func (r *RedisNotificationRepository) GetTokensByUserIDs(ctx context.Context, userIDs []string) (map[string]string, error) {
	result := make(map[string]string)
	for _, uid := range userIDs {
		val, err := r.rdb.HGet(ctx, "fcm:tokens", uid).Result()
		if err == redis.Nil {
			continue
		}
		if err != nil {
			return nil, err
		}
		result[uid] = val
	}
	return result, nil
}
