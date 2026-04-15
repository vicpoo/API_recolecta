package application

import (
	"context"
	"fmt"
	"strings"

	"github.com/vicpoo/API_recolecta/src/notificacion/domain"
)

type QueryEventTraceUseCase struct {
	repo domain.EventTraceRepository
}

func NewQueryEventTraceUseCase(repo domain.EventTraceRepository) *QueryEventTraceUseCase {
	return &QueryEventTraceUseCase{repo: repo}
}

func (uc *QueryEventTraceUseCase) GetByEventID(ctx context.Context, eventID string) (*domain.EventTraceRecord, error) {
	normalized := strings.TrimSpace(eventID)
	if normalized == "" {
		return nil, fmt.Errorf("event_id es obligatorio")
	}
	return uc.repo.GetByEventID(ctx, normalized)
}

func (uc *QueryEventTraceUseCase) ListByTruckID(ctx context.Context, truckID int32, limit int64) ([]domain.EventTraceRecord, error) {
	if truckID <= 0 {
		return nil, fmt.Errorf("truck_id invalido")
	}
	return uc.repo.ListByTruckID(ctx, truckID, limit)
}
