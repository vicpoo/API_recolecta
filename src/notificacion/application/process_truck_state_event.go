package application

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"fmt"

	"github.com/vicpoo/API_recolecta/src/notificacion/domain"
)

type ProcessTruckStateEventUseCase struct {
	rulesRepo domain.INotificationRuleRepository
	traceRepo domain.IEventTraceRepository
}

func NewProcessTruckStateEventUseCase(rulesRepo domain.INotificationRuleRepository, traceRepo domain.IEventTraceRepository) *ProcessTruckStateEventUseCase {
	return &ProcessTruckStateEventUseCase{rulesRepo: rulesRepo, traceRepo: traceRepo}
}

func (uc *ProcessTruckStateEventUseCase) Execute(ctx context.Context, event *domain.TruckStateEvent) error {
	hash := computeEventHash(event)

	acquired, err := uc.traceRepo.TryAcquireDeduplication(ctx, hash, event)
	if err != nil {
		return err
	}
	if !acquired {
		return nil
	}

	action := domain.ActionNotifyAdminOnly
	adminNotified := true

	rule, err := uc.rulesRepo.GetByStateCode(ctx, string(event.StateCode))
	if err == nil && !rule.Enabled {
		adminNotified = false
	}

	trace := &domain.EventTraceRecord{
		EventID:        event.EventID,
		EventHash:      hash,
		EventType:      event.EventType,
		EventVersion:   event.EventVersion,
		TruckID:        event.TruckID,
		StateCode:      event.StateCode,
		ResolvedAction: action,
		AdminNotified:  adminNotified,
		Result:         "processed",
		CreatedAt:      event.OccurredAt,
	}
	return uc.traceRepo.SaveTrace(ctx, trace)
}

func computeEventHash(event *domain.TruckStateEvent) string {
	raw := fmt.Sprintf("%s:%d:%s:%s", event.EventID, event.TruckID, event.StateCode, event.EventType)
	h := sha256.Sum256([]byte(raw))
	return hex.EncodeToString(h[:])
}
