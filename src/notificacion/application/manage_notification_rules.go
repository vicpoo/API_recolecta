package application

import (
	"context"
	"fmt"
	"strings"

	"github.com/vicpoo/API_recolecta/src/notificacion/domain"
)

type ManageNotificationRulesUseCase struct {
	repo domain.NotificationRuleRepository
}

func NewManageNotificationRulesUseCase(repo domain.NotificationRuleRepository) *ManageNotificationRulesUseCase {
	return &ManageNotificationRulesUseCase{repo: repo}
}

func (uc *ManageNotificationRulesUseCase) Upsert(ctx context.Context, rule *domain.NotificationRule) error {
	stateCode := strings.TrimSpace(strings.ToUpper(rule.StateCode))
	if stateCode == "" {
		return fmt.Errorf("state_code es obligatorio")
	}
	if strings.TrimSpace(rule.Action) == "" {
		return fmt.Errorf("action es obligatorio")
	}
	if rule.RadiusMeters < 0 {
		return fmt.Errorf("radius_meters no puede ser negativo")
	}

	rule.StateCode = stateCode
	return uc.repo.Upsert(ctx, rule)
}

func (uc *ManageNotificationRulesUseCase) GetByStateCode(ctx context.Context, stateCode string) (*domain.NotificationRule, error) {
	normalized := strings.TrimSpace(strings.ToUpper(stateCode))
	if normalized == "" {
		return nil, fmt.Errorf("state_code es obligatorio")
	}
	return uc.repo.GetByStateCode(ctx, normalized)
}

func (uc *ManageNotificationRulesUseCase) List(ctx context.Context) ([]domain.NotificationRule, error) {
	return uc.repo.List(ctx)
}

func (uc *ManageNotificationRulesUseCase) Delete(ctx context.Context, stateCode string) error {
	normalized := strings.TrimSpace(strings.ToUpper(stateCode))
	if normalized == "" {
		return fmt.Errorf("state_code es obligatorio")
	}
	return uc.repo.Delete(ctx, normalized)
}
