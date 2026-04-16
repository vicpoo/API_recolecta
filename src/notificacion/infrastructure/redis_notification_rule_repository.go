package infrastructure

import (
	"context"
	"fmt"
	"sort"
	"strconv"
	"strings"
	"time"

	goredis "github.com/redis/go-redis/v9"
	"github.com/vicpoo/API_recolecta/src/notificacion/domain"
)

const (
	notificationRuleKeyFmt = "rules:state:%s"
	notificationRuleScan   = "rules:state:*"
	notificationRulesVer   = "rules:version"
)

type RedisNotificationRuleRepository struct {
	rdb *goredis.Client
}

func NewRedisNotificationRuleRepository(rdb *goredis.Client) *RedisNotificationRuleRepository {
	return &RedisNotificationRuleRepository{rdb: rdb}
}

func (r *RedisNotificationRuleRepository) Upsert(ctx context.Context, rule *domain.NotificationRule) error {
	version, err := r.rdb.Incr(ctx, notificationRulesVer).Result()
	if err != nil {
		return fmt.Errorf("no se pudo incrementar rules:version: %w", err)
	}

	key := fmt.Sprintf(notificationRuleKeyFmt, strings.ToUpper(rule.StateCode))
	payload := map[string]interface{}{
		"state_code":     strings.ToUpper(rule.StateCode),
		"action":         rule.Action,
		"radius_meters":  rule.RadiusMeters,
		"priority":       rule.Priority,
		"enabled":        strconv.FormatBool(rule.Enabled),
		"template_title": rule.TemplateTitle,
		"template_body":  rule.TemplateBody,
		"version":        version,
		"updated_at":     time.Now().UTC().Format(time.RFC3339),
	}

	if err := r.rdb.HSet(ctx, key, payload).Err(); err != nil {
		return fmt.Errorf("no se pudo guardar regla %s: %w", rule.StateCode, err)
	}

	rule.Version = version
	return nil
}

func (r *RedisNotificationRuleRepository) GetByStateCode(ctx context.Context, stateCode string) (*domain.NotificationRule, error) {
	key := fmt.Sprintf(notificationRuleKeyFmt, strings.ToUpper(stateCode))
	values, err := r.rdb.HGetAll(ctx, key).Result()
	if err != nil {
		return nil, fmt.Errorf("no se pudo leer regla %s: %w", stateCode, err)
	}
	if len(values) == 0 {
		return nil, fmt.Errorf("no existe regla para state_code %s", stateCode)
	}

	rule, err := mapToNotificationRule(values)
	if err != nil {
		return nil, err
	}
	return &rule, nil
}

func (r *RedisNotificationRuleRepository) List(ctx context.Context) ([]domain.NotificationRule, error) {
	keys := make([]string, 0)
	iter := r.rdb.Scan(ctx, 0, notificationRuleScan, 100).Iterator()
	for iter.Next(ctx) {
		keys = append(keys, iter.Val())
	}
	if err := iter.Err(); err != nil {
		return nil, fmt.Errorf("no se pudieron listar reglas: %w", err)
	}
	if len(keys) == 0 {
		return []domain.NotificationRule{}, nil
	}

	values, err := r.rdb.Pipelined(ctx, func(pipe goredis.Pipeliner) error {
		for _, key := range keys {
			pipe.HGetAll(ctx, key)
		}
		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("no se pudieron leer reglas: %w", err)
	}

	rules := make([]domain.NotificationRule, 0, len(values))
	for _, cmd := range values {
		mapCmd, ok := cmd.(*goredis.MapStringStringCmd)
		if !ok {
			continue
		}
		ruleMap, cmdErr := mapCmd.Result()
		if cmdErr != nil || len(ruleMap) == 0 {
			continue
		}
		rule, parseErr := mapToNotificationRule(ruleMap)
		if parseErr != nil {
			return nil, parseErr
		}
		rules = append(rules, rule)
	}

	sort.Slice(rules, func(i, j int) bool {
		return rules[i].StateCode < rules[j].StateCode
	})
	return rules, nil
}

func (r *RedisNotificationRuleRepository) Delete(ctx context.Context, stateCode string) error {
	if err := r.rdb.Del(ctx, fmt.Sprintf(notificationRuleKeyFmt, strings.ToUpper(stateCode))).Err(); err != nil {
		return fmt.Errorf("no se pudo eliminar regla %s: %w", stateCode, err)
	}
	if _, err := r.rdb.Incr(ctx, notificationRulesVer).Result(); err != nil {
		return fmt.Errorf("no se pudo incrementar rules:version tras delete: %w", err)
	}
	return nil
}

func mapToNotificationRule(values map[string]string) (domain.NotificationRule, error) {
	radius, err := strconv.Atoi(values["radius_meters"])
	if err != nil {
		return domain.NotificationRule{}, fmt.Errorf("radius_meters inválido: %w", err)
	}
	priority, err := strconv.Atoi(values["priority"])
	if err != nil {
		return domain.NotificationRule{}, fmt.Errorf("priority inválido: %w", err)
	}
	enabled, err := strconv.ParseBool(values["enabled"])
	if err != nil {
		return domain.NotificationRule{}, fmt.Errorf("enabled inválido: %w", err)
	}
	version, err := strconv.ParseInt(values["version"], 10, 64)
	if err != nil {
		return domain.NotificationRule{}, fmt.Errorf("version inválido: %w", err)
	}

	return domain.NotificationRule{
		StateCode:     values["state_code"],
		Action:        values["action"],
		RadiusMeters:  radius,
		Priority:      priority,
		Enabled:       enabled,
		TemplateTitle: values["template_title"],
		TemplateBody:  values["template_body"],
		Version:       version,
	}, nil
}
