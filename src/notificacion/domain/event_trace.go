package domain

import (
	"context"
	"time"
)

type EventTraceRecord struct {
	EventID            string    `json:"event_id"`
	EventHash          string    `json:"event_hash"`
	EventType          string    `json:"event_type"`
	EventVersion       string    `json:"event_version"`
	TruckID            int32     `json:"truck_id"`
	StateCode          string    `json:"state_code"`
	ResolvedAction     string    `json:"resolved_action"`
	AdminNotified      bool      `json:"admin_notified"`
	CitizenFanoutCount int       `json:"citizen_fanout_count"`
	Result             string    `json:"result"`
	CreatedAt          time.Time `json:"created_at"`
}

type EventTraceRepository interface {
	TryAcquireDeduplication(ctx context.Context, eventHash string, event *TruckStateEvent) (bool, error)
	SaveTrace(ctx context.Context, trace *EventTraceRecord) error
	GetByEventID(ctx context.Context, eventID string) (*EventTraceRecord, error)
	ListByTruckID(ctx context.Context, truckID int32, limit int64) ([]EventTraceRecord, error)
	CountByTruckID(ctx context.Context, truckID int32) (int64, error)
}
