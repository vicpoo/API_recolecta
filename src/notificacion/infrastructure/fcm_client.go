package infrastructure

import (
	"context"
	"fmt"
	"os"
	"sort"
	"strings"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/messaging"
	"github.com/vicpoo/API_recolecta/src/notificacion/domain"
	"google.golang.org/api/option"
)

type FCMClient struct {
	client *messaging.Client
}

func NewFCMClient() (*FCMClient, error) {
	credentialsPath := resolveCredentialsPath()
	if credentialsPath == "" {
		return nil, fmt.Errorf("no se definio archivo de credenciales FCM (FCM_CREDENTIALS_FILE o GOOGLE_APPLICATION_CREDENTIALS)")
	}
	if _, err := os.Stat(credentialsPath); err != nil {
		return nil, fmt.Errorf("archivo de credenciales FCM no accesible en %s: %w", credentialsPath, err)
	}

	cfg := &firebase.Config{}
	if projectID := resolveProjectID(); projectID != "" {
		cfg.ProjectID = projectID
	}

	app, err := firebase.NewApp(context.Background(), cfg, option.WithCredentialsFile(credentialsPath))
	if err != nil {
		return nil, fmt.Errorf("error initializing firebase app: %w", err)
	}

	info, err := os.Stat(credentialPath)
	if err != nil {
		return nil, fmt.Errorf("fcm credentials file not found at '%s': %w", credentialPath, err)
	}
	if info.IsDir() {
		return nil, fmt.Errorf("fcm credentials path points to a directory, expected file: '%s'", credentialPath)
	}

	app, err := firebase.NewApp(ctx, nil, option.WithCredentialsFile(credentialPath))
	if err != nil {
		return nil, fmt.Errorf("error initializing firebase app with credentials file '%s': %w", credentialPath, err)
	}

	client, err := app.Messaging(ctx)
	if err != nil {
		return nil, fmt.Errorf("error creating firebase messaging client: %w", err)
	}

	return &FCMClient{client: client}, nil
}

func resolveCredentialsPath() string {
	candidates := []string{
		strings.TrimSpace(os.Getenv("FCM_CREDENTIALS_FILE")),
		strings.TrimSpace(os.Getenv("GOOGLE_APPLICATION_CREDENTIALS")),
	}
	for _, candidate := range candidates {
		if candidate != "" {
			return candidate
		}
	}
	return ""
}

func resolveProjectID() string {
	candidates := []string{
		strings.TrimSpace(os.Getenv("FIREBASE_PROJECT_ID")),
		strings.TrimSpace(os.Getenv("GOOGLE_CLOUD_PROJECT")),
		strings.TrimSpace(os.Getenv("GCLOUD_PROJECT")),
	}
	for _, candidate := range candidates {
		if candidate != "" {
			return candidate
		}
	}
	return ""
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
