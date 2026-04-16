package application

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/vicpoo/API_recolecta/src/notificacion/domain"
)

type SendCitizenNotificationInput struct {
	UserTokens map[string]string `json:"user_tokens"`
	Title      string            `json:"title"`
	Body       string            `json:"body"`
	Type       string            `json:"type"`
	Data       map[string]string `json:"data,omitempty"`
}

type SendCitizenNotificationOutput struct {
	Results map[string]domain.SendResult `json:"results"`
	Error   error                        `json:"-"`
}

type SendCitizenNotificationUseCase struct {
	notificationSender domain.NotificationSender
	logRepo            domain.NotificationLogRepository
}

func NewSendCitizenNotificationUseCase(
	sender domain.NotificationSender,
	repo domain.NotificationLogRepository,
) *SendCitizenNotificationUseCase {
	return &SendCitizenNotificationUseCase{
		notificationSender: sender,
		logRepo:            repo,
	}
}

func (uc *SendCitizenNotificationUseCase) Execute(ctx context.Context, input *SendCitizenNotificationInput) *SendCitizenNotificationOutput {
	output := &SendCitizenNotificationOutput{
		Results: make(map[string]domain.SendResult),
	}

	if len(input.UserTokens) == 0 {
		output.Error = fmt.Errorf("no hay tokens de usuario para enviar")
		return output
	}

	if input.Title == "" || input.Body == "" {
		output.Error = fmt.Errorf("titulo o cuerpo de la notificacion vacios")
		return output
	}

	validUserTokens := make(map[string]string, len(input.UserTokens))
	notificationIDs := make(map[string]string, len(input.UserTokens))
	currentTime := time.Now().UTC()

	for userID, token := range input.UserTokens {
		if token == "" {
			output.Results[userID] = domain.SendResult{
				Success: false,
				Error:   "token FCM vacio",
			}
			continue
		}

		notificationID := uuid.NewString()
		notificationIDs[userID] = notificationID

		entry := &domain.NotificationLogEntry{
			NotificationID: notificationID,
			UserID:         userID,
			Title:          input.Title,
			Body:           input.Body,
			Type:           input.Type,
			Status:         "pending",
			Timestamp:      currentTime,
		}

		if err := uc.logRepo.CreateNotificationLogEntry(ctx, entry); err != nil {
			output.Results[userID] = domain.SendResult{
				Success: false,
				Error:   fmt.Sprintf("error interno al registrar log: %v", err),
			}
			continue
		}

		if err := uc.logRepo.AddNotificationToUserLog(ctx, userID, notificationID, currentTime); err != nil {
			output.Results[userID] = domain.SendResult{
				Success: false,
				Error:   fmt.Sprintf("error interno al registrar log de usuario: %v", err),
			}
			_ = uc.logRepo.UpdateNotificationStatus(ctx, notificationID, "failed", err.Error())
			continue
		}

		validUserTokens[userID] = token
	}

	if len(validUserTokens) == 0 {
		output.Error = fmt.Errorf("no quedan tokens validos despues del pre-registro")
		return output
	}

	notificationPayload := &domain.PushNotification{
		Title: input.Title,
		Body:  input.Body,
		Type:  input.Type,
		Data:  input.Data,
	}

	sendResults, err := uc.notificationSender.Send(ctx, validUserTokens, notificationPayload)
	if err != nil {
		output.Error = fmt.Errorf("error global al enviar notificaciones FCM: %v", err)
		for userID, notificationID := range notificationIDs {
			if _, alreadyProcessed := output.Results[userID]; alreadyProcessed {
				continue
			}
			_ = uc.logRepo.UpdateNotificationStatus(ctx, notificationID, "failed", err.Error())
			output.Results[userID] = domain.SendResult{Success: false, Error: err.Error()}
		}
		return output
	}

	for userID, result := range sendResults {
		notificationID, ok := notificationIDs[userID]
		if !ok {
			continue
		}

		if result.Success {
			_ = uc.logRepo.UpdateNotificationStatus(ctx, notificationID, "delivered", "")
		} else {
			_ = uc.logRepo.UpdateNotificationStatus(ctx, notificationID, "failed", result.Error)
		}

		output.Results[userID] = result
	}

	return output
}