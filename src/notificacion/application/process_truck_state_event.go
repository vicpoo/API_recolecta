package application

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/vicpoo/API_recolecta/src/notificacion/domain"
)

type ProcessTruckStateEventOutput struct {
	EventID        string `json:"event_id"`
	EventHash      string `json:"event_hash"`
	ResolvedAction string `json:"resolved_action"`
	Deduplicated   bool   `json:"deduplicated"`
	Result         string `json:"result"`
}

type ProcessTruckStateEventUseCase struct {
	rulesRepo domain.NotificationRuleRepository
	traceRepo domain.EventTraceRepository
}

func NewProcessTruckStateEventUseCase(
	rulesRepo domain.NotificationRuleRepository,
	traceRepo domain.EventTraceRepository,
) *ProcessTruckStateEventUseCase {
	return &ProcessTruckStateEventUseCase{
		rulesRepo: rulesRepo,
		traceRepo: traceRepo,
	}
}

func (uc *ProcessTruckStateEventUseCase) Execute(ctx context.Context, event *domain.TruckStateEvent) (*ProcessTruckStateEventOutput, error) {
	if strings.TrimSpace(event.EventID) == "" {
		return nil, fmt.Errorf("event_id es obligatorio")
	}
	if strings.TrimSpace(event.EventType) == "" {
		return nil, fmt.Errorf("event_type es obligatorio")
	}
	if event.TruckID <= 0 {
		return nil, fmt.Errorf("truck_id es obligatorio")
	}
	if event.OccurredAt.IsZero() {
		return nil, fmt.Errorf("occurred_at es obligatorio")
	}
	if event.EventVersion != domain.EventVersionV1 {
		return nil, fmt.Errorf("event_version invalido: %s", event.EventVersion)
	}

	stateCode := strings.ToUpper(strings.TrimSpace(fmt.Sprint(event.Payload["state_code"])))
	if stateCode == "" || stateCode == "<nil>" {
		return nil, fmt.Errorf("payload.state_code es obligatorio")
	}

	eventHash := buildEventHash(event, stateCode)
	acquired, err := uc.traceRepo.TryAcquireDeduplication(ctx, eventHash, event)
	if err != nil {
		return nil, err
	}
	if !acquired {
		return &ProcessTruckStateEventOutput{
			EventID:        event.EventID,
			EventHash:      eventHash,
			ResolvedAction: "",
			Deduplicated:   true,
			Result:         "duplicate",
		}, nil
	}

	rule, err := uc.rulesRepo.GetByStateCode(ctx, stateCode)
	if err != nil {
		return nil, err
	}

	resolvedAction := rule.Action
	adminNotified := false
	citizenFanout := 0
	result := "processed"

	if !rule.Enabled {
		resolvedAction = "DISABLED"
		result = "rule_disabled"
	} else {
		switch rule.Action {
		case domain.ActionNotifyAdminOnly:
			adminNotified = true
		case domain.ActionNotifyAdminAndCitizens:
			adminNotified = true
			citizenFanout = extractCitizenFanout(event.Payload["citizen_fanout_count"])
		}
	}

	trace := &domain.EventTraceRecord{
		EventID:            event.EventID,
		EventHash:          eventHash,
		EventType:          event.EventType,
		EventVersion:       event.EventVersion,
		TruckID:            event.TruckID,
		StateCode:          stateCode,
		ResolvedAction:     resolvedAction,
		AdminNotified:      adminNotified,
		CitizenFanoutCount: citizenFanout,
		Result:             result,
		CreatedAt:          time.Now().UTC(),
	}

	if err := uc.traceRepo.SaveTrace(ctx, trace); err != nil {
		return nil, err
	}

	return &ProcessTruckStateEventOutput{
		EventID:        event.EventID,
		EventHash:      eventHash,
		ResolvedAction: resolvedAction,
		Deduplicated:   false,
		Result:         result,
	}, nil
}

func buildEventHash(event *domain.TruckStateEvent, stateCode string) string {
	pointID := strings.TrimSpace(fmt.Sprint(event.Payload["point_id"]))
	hashInput := strings.Join([]string{
		event.EventType,
		event.EventVersion,
		strconv.FormatInt(int64(event.TruckID), 10),
		stateCode,
		pointID,
		event.OccurredAt.UTC().Format(time.RFC3339Nano),
	}, "|")
	hash := sha256.Sum256([]byte(hashInput))
	return hex.EncodeToString(hash[:])
}

func extractCitizenFanout(value interface{}) int {
	switch typed := value.(type) {
	case int:
		if typed < 0 {
			return 0
		}
		return typed
	case int32:
		if typed < 0 {
			return 0
		}
		return int(typed)
	case int64:
		if typed < 0 {
			return 0
		}
		return int(typed)
	case float64:
		if typed < 0 {
			return 0
		}
		return int(typed)
	case string:
		parsed, err := strconv.Atoi(strings.TrimSpace(typed))
		if err != nil || parsed < 0 {
			return 0
		}
		return parsed
	default:
		return 0
	}
}
