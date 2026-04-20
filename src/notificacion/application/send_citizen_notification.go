package application

import (
	"context"

	"github.com/vicpoo/API_recolecta/src/notificacion/domain"
)

type tokenRepository interface {
	GetTokensByUserIDs(ctx context.Context, userIDs []string) (map[string]string, error)
}

type SendCitizenNotificationUseCase struct {
	sender domain.PushNotificationSender
	repo   tokenRepository
}

func NewSendCitizenNotificationUseCase(sender domain.PushNotificationSender, repo tokenRepository) *SendCitizenNotificationUseCase {
	return &SendCitizenNotificationUseCase{sender: sender, repo: repo}
}

func (uc *SendCitizenNotificationUseCase) Execute(ctx context.Context, userIDs []string, notification *domain.PushNotification) (map[string]domain.SendResult, error) {
	tokens, err := uc.repo.GetTokensByUserIDs(ctx, userIDs)
	if err != nil {
		return nil, err
	}
	return uc.sender.Send(ctx, tokens, notification)
}
