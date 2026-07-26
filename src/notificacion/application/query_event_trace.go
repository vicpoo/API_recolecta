package application

import (
	"context"

	"github.com/vicpoo/API_recolecta/src/notificacion/domain"
)

type QueryEventTraceUseCase struct {
	repo domain.IEventTraceRepository
}

func NewQueryEventTraceUseCase(repo domain.IEventTraceRepository) *QueryEventTraceUseCase {
	return &QueryEventTraceUseCase{repo: repo}
}

func (uc *QueryEventTraceUseCase) GetByEventID(ctx context.Context, eventID string) (*domain.EventTraceRecord, error) {
	return uc.repo.GetByEventID(ctx, eventID)
}

func (uc *QueryEventTraceUseCase) ListByTruckID(ctx context.Context, truckID int32, limit int) ([]domain.EventTraceRecord, error) {
	return uc.repo.ListByTruckID(ctx, truckID, limit)
}
