package domain

import (
	"context"
	"time"
)

type SendResult struct {
	Success bool   `json:"success"`
	Error   string `json:"error,omitempty"`
}

type PushNotification struct {
	Title string            `json:"title"`
	Body  string            `json:"body"`
	Type  string            `json:"type"`
	Data  map[string]string `json:"data,omitempty"`
}

type NotificationSender interface {
	Send(ctx context.Context, userTokens map[string]string, notification *PushNotification) (map[string]SendResult, error)
}

type NotificationLogEntry struct {
	NotificationID string    `json:"notification_id"`
	UserID         string    `json:"user_id"`
	Title          string    `json:"title"`
	Body           string    `json:"body"`
	Type           string    `json:"type"`
	Status         string    `json:"status"`
	ErrorMessage   string    `json:"error_message,omitempty"`
	Timestamp      time.Time `json:"timestamp"`
}

type NotificationLogRepository interface {
	CreateNotificationLogEntry(ctx context.Context, entry *NotificationLogEntry) error
	UpdateNotificationStatus(ctx context.Context, notificationID string, status string, errorMessage string) error
	AddNotificationToUserLog(ctx context.Context, userID string, notificationID string, timestamp time.Time) error
}