package domain

import "context"

type NotificationRule struct {
	StateCode string `json:"state_code"`
	Enabled   bool   `json:"enabled"`
	Threshold int    `json:"threshold"`
}

type INotificationRuleRepository interface {
	Save(ctx context.Context, rule NotificationRule) error
	GetByStateCode(ctx context.Context, code string) (*NotificationRule, error)
	List(ctx context.Context) ([]NotificationRule, error)
	Delete(ctx context.Context, code string) error
}
