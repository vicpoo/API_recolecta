package domain

import "context"

type PushNotification struct {
	Title string
	Body  string
	Type  string
	Data  map[string]string
}

type SendResult struct {
	Success bool
	Error   string
}

type PushNotificationSender interface {
	Send(ctx context.Context, userTokens map[string]string, notification *PushNotification) (map[string]SendResult, error)
}
