package infrastructure

import (
	"context"
	"fmt"
	"sort"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/messaging"
	"github.com/vicpoo/API_recolecta/src/notificacion/domain"
)

type FCMClient struct {
	client *messaging.Client
}

func NewFCMClient() (*FCMClient, error) {
	app, err := firebase.NewApp(context.Background(), nil)
	if err != nil {
		return nil, fmt.Errorf("error initializing firebase app via ADC: %w", err)
	}

	client, err := app.Messaging(context.Background())
	if err != nil {
		return nil, fmt.Errorf("error creating firebase messaging client: %w", err)
	}

	return &FCMClient{client: client}, nil
}

func (c *FCMClient) Send(ctx context.Context, userTokens map[string]string, notification *domain.PushNotification) (map[string]domain.SendResult, error) {
	userIDs := make([]string, 0, len(userTokens))
	for userID := range userTokens {
		userIDs = append(userIDs, userID)
	}
	sort.Strings(userIDs)

	tokens := make([]string, 0, len(userIDs))
	for _, userID := range userIDs {
		tokens = append(tokens, userTokens[userID])
	}

	dataPayload := make(map[string]string, len(notification.Data)+1)
	for key, value := range notification.Data {
		dataPayload[key] = value
	}
	if notification.Type != "" {
		dataPayload["notificationType"] = notification.Type
	}

	message := &messaging.MulticastMessage{
		Tokens: tokens,
		Notification: &messaging.Notification{
			Title: notification.Title,
			Body:  notification.Body,
		},
		Data: dataPayload,
	}

	batchResponse, err := c.client.SendEachForMulticast(ctx, message)
	if err != nil {
		return nil, err
	}

	results := make(map[string]domain.SendResult, len(userIDs))
	for index, response := range batchResponse.Responses {
		userID := userIDs[index]
		if response.Success {
			results[userID] = domain.SendResult{Success: true}
			continue
		}

		errorMessage := "unknown error"
		if response.Error != nil {
			errorMessage = response.Error.Error()
		}

		results[userID] = domain.SendResult{
			Success: false,
			Error:   errorMessage,
		}
	}

	return results, nil
}