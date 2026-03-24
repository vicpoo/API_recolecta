package infrastructure

import (
	"context"
	"fmt"
	"time"

	goredis "github.com/redis/go-redis/v9"
	"github.com/vicpoo/API_recolecta/src/notificacion/domain"
)

const (
	notificationHashKeyFmt   = "notification:%s"
	userNotificationLogKeyFmt = "notification:log:%s"
)

type RedisNotificationRepository struct {
	rdb *goredis.Client
}

func NewRedisNotificationRepository(rdb *goredis.Client) *RedisNotificationRepository {
	return &RedisNotificationRepository{rdb: rdb}
}

func (r *RedisNotificationRepository) CreateNotificationLogEntry(ctx context.Context, entry *domain.NotificationLogEntry) error {
	return r.rdb.HSet(ctx, fmt.Sprintf(notificationHashKeyFmt, entry.NotificationID),
		"notification_id", entry.NotificationID,
		"user_id", entry.UserID,
		"title", entry.Title,
		"body", entry.Body,
		"type", entry.Type,
		"status", entry.Status,
		"error_message", entry.ErrorMessage,
		"timestamp", entry.Timestamp.Format(time.RFC3339),
	).Err()
}

func (r *RedisNotificationRepository) UpdateNotificationStatus(ctx context.Context, notificationID string, status string, errorMessage string) error {
	return r.rdb.HSet(ctx, fmt.Sprintf(notificationHashKeyFmt, notificationID),
		"status", status,
		"error_message", errorMessage,
	).Err()
}

func (r *RedisNotificationRepository) AddNotificationToUserLog(ctx context.Context, userID string, notificationID string, timestamp time.Time) error {
	return r.rdb.ZAdd(ctx, fmt.Sprintf(userNotificationLogKeyFmt, userID), goredis.Z{
		Score:  float64(timestamp.Unix()),
		Member: notificationID,
	}).Err()
}