package domain

import "context"

type NotificationRule struct {
	StateCode     string `json:"state_code"`
	Action        string `json:"action"`
	RadiusMeters  int    `json:"radius_meters"`
	Priority      int    `json:"priority"`
	Enabled       bool   `json:"enabled"`
	TemplateTitle string `json:"template_title"`
	TemplateBody  string `json:"template_body"`
	Version       int64  `json:"version"`
}

type NotificationRuleRepository interface {
	Upsert(ctx context.Context, rule *NotificationRule) error
	GetByStateCode(ctx context.Context, stateCode string) (*NotificationRule, error)
	List(ctx context.Context) ([]NotificationRule, error)
	Delete(ctx context.Context, stateCode string) error
}
