package application

import (
	"context"
	"fmt"

	"github.com/vicpoo/API_recolecta/src/notificacion/domain"
)

type ObservabilitySummaryOutput struct {
	TruckID          int32 `json:"truck_id"`
	TotalTruckTraces int64 `json:"total_truck_traces"`
	ActiveWSSessions int64 `json:"active_ws_sessions"`
}

type GetObservabilitySummaryUseCase struct {
	traceRepo   domain.EventTraceRepository
	sessionRepo domain.AdminRealtimeSessionRepository
}

func NewGetObservabilitySummaryUseCase(
	traceRepo domain.EventTraceRepository,
	sessionRepo domain.AdminRealtimeSessionRepository,
) *GetObservabilitySummaryUseCase {
	return &GetObservabilitySummaryUseCase{
		traceRepo:   traceRepo,
		sessionRepo: sessionRepo,
	}
}

func (uc *GetObservabilitySummaryUseCase) Execute(ctx context.Context, truckID int32) (*ObservabilitySummaryOutput, error) {
	if truckID <= 0 {
		return nil, fmt.Errorf("truck_id invalido")
	}

	totalTruckTraces, err := uc.traceRepo.CountByTruckID(ctx, truckID)
	if err != nil {
		return nil, err
	}
	activeSessions, err := uc.sessionRepo.CountActiveSessions(ctx)
	if err != nil {
		return nil, err
	}

	return &ObservabilitySummaryOutput{
		TruckID:          truckID,
		TotalTruckTraces: totalTruckTraces,
		ActiveWSSessions: activeSessions,
	}, nil
}
