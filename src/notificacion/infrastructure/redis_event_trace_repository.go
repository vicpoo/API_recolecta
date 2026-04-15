package infrastructure

import (
	"context"
	"fmt"
	"strconv"
	"time"

	goredis "github.com/redis/go-redis/v9"
	"github.com/vicpoo/API_recolecta/src/notificacion/domain"
)

const (
	eventDeduplicationKeyFmt = "event_deduplication:%s"
	eventTraceKeyFmt         = "event_trace:%s"
	eventTraceTruckKeyFmt    = "event_trace:truck:%d"
	eventRetentionTTL        = 30 * 24 * time.Hour
)

type RedisEventTraceRepository struct {
	rdb *goredis.Client
}

func NewRedisEventTraceRepository(rdb *goredis.Client) *RedisEventTraceRepository {
	return &RedisEventTraceRepository{rdb: rdb}
}

func (r *RedisEventTraceRepository) TryAcquireDeduplication(ctx context.Context, eventHash string, event *domain.TruckStateEvent) (bool, error) {
	key := fmt.Sprintf(eventDeduplicationKeyFmt, eventHash)
	createdAt := time.Now().UTC()

	res, err := r.rdb.TxPipelined(ctx, func(pipe goredis.Pipeliner) error {
		pipe.HSetNX(ctx, key, "event_id", event.EventID)
		pipe.HSetNX(ctx, key, "event_type", event.EventType)
		pipe.HSetNX(ctx, key, "truck_id", event.TruckID)
		pipe.HSetNX(ctx, key, "processed_at", createdAt.Format(time.RFC3339))
		pipe.HSetNX(ctx, key, "result", "processed")
		pipe.Expire(ctx, key, eventRetentionTTL)
		return nil
	})
	if err != nil {
		return false, fmt.Errorf("no se pudo registrar dedupe para %s: %w", eventHash, err)
	}

	if len(res) == 0 {
		return false, fmt.Errorf("resultado vacio al registrar dedupe")
	}

	firstCmd, ok := res[0].(*goredis.BoolCmd)
	if !ok {
		return false, fmt.Errorf("resultado inesperado al registrar dedupe")
	}

	acquired, err := firstCmd.Result()
	if err != nil {
		return false, fmt.Errorf("no se pudo evaluar dedupe para %s: %w", eventHash, err)
	}
	return acquired, nil
}

func (r *RedisEventTraceRepository) SaveTrace(ctx context.Context, trace *domain.EventTraceRecord) error {
	traceKey := fmt.Sprintf(eventTraceKeyFmt, trace.EventID)
	truckTraceKey := fmt.Sprintf(eventTraceTruckKeyFmt, trace.TruckID)

	_, err := r.rdb.TxPipelined(ctx, func(pipe goredis.Pipeliner) error {
		pipe.HSet(ctx, traceKey,
			"event_id", trace.EventID,
			"event_hash", trace.EventHash,
			"event_type", trace.EventType,
			"event_version", trace.EventVersion,
			"truck_id", trace.TruckID,
			"state_code", trace.StateCode,
			"resolved_action", trace.ResolvedAction,
			"admin_notified", strconv.FormatBool(trace.AdminNotified),
			"citizen_fanout_count", trace.CitizenFanoutCount,
			"result", trace.Result,
			"created_at", trace.CreatedAt.UTC().Format(time.RFC3339),
		)
		pipe.Expire(ctx, traceKey, eventRetentionTTL)
		pipe.ZAdd(ctx, truckTraceKey, goredis.Z{
			Score:  float64(trace.CreatedAt.Unix()),
			Member: trace.EventID,
		})
		pipe.Expire(ctx, truckTraceKey, eventRetentionTTL)
		return nil
	})
	if err != nil {
		return fmt.Errorf("no se pudo guardar traza de evento %s: %w", trace.EventID, err)
	}

	return nil
}

func (r *RedisEventTraceRepository) GetByEventID(ctx context.Context, eventID string) (*domain.EventTraceRecord, error) {
	values, err := r.rdb.HGetAll(ctx, fmt.Sprintf(eventTraceKeyFmt, eventID)).Result()
	if err != nil {
		return nil, fmt.Errorf("no se pudo leer traza %s: %w", eventID, err)
	}
	if len(values) == 0 {
		return nil, fmt.Errorf("traza no encontrada para event_id %s", eventID)
	}
	trace, err := mapToEventTrace(values)
	if err != nil {
		return nil, err
	}
	return &trace, nil
}

func (r *RedisEventTraceRepository) ListByTruckID(ctx context.Context, truckID int32, limit int64) ([]domain.EventTraceRecord, error) {
	if limit <= 0 {
		limit = 20
	}
	indexKey := fmt.Sprintf(eventTraceTruckKeyFmt, truckID)
	eventIDs, err := r.rdb.ZRevRange(ctx, indexKey, 0, limit-1).Result()
	if err != nil {
		return nil, fmt.Errorf("no se pudo listar trazas de truck_id %d: %w", truckID, err)
	}
	if len(eventIDs) == 0 {
		return []domain.EventTraceRecord{}, nil
	}

	traces := make([]domain.EventTraceRecord, 0, len(eventIDs))
	for _, eventID := range eventIDs {
		trace, getErr := r.GetByEventID(ctx, eventID)
		if getErr != nil {
			return nil, getErr
		}
		traces = append(traces, *trace)
	}
	return traces, nil
}

func (r *RedisEventTraceRepository) CountByTruckID(ctx context.Context, truckID int32) (int64, error) {
	total, err := r.rdb.ZCard(ctx, fmt.Sprintf(eventTraceTruckKeyFmt, truckID)).Result()
	if err != nil {
		return 0, fmt.Errorf("no se pudo contar trazas de truck_id %d: %w", truckID, err)
	}
	return total, nil
}

func mapToEventTrace(values map[string]string) (domain.EventTraceRecord, error) {
	truckID, err := strconv.Atoi(values["truck_id"])
	if err != nil {
		return domain.EventTraceRecord{}, fmt.Errorf("truck_id invalido en traza: %w", err)
	}
	adminNotified, err := strconv.ParseBool(values["admin_notified"])
	if err != nil {
		return domain.EventTraceRecord{}, fmt.Errorf("admin_notified invalido en traza: %w", err)
	}
	citizenFanoutCount, err := strconv.Atoi(values["citizen_fanout_count"])
	if err != nil {
		return domain.EventTraceRecord{}, fmt.Errorf("citizen_fanout_count invalido en traza: %w", err)
	}
	createdAt, err := time.Parse(time.RFC3339, values["created_at"])
	if err != nil {
		return domain.EventTraceRecord{}, fmt.Errorf("created_at invalido en traza: %w", err)
	}

	return domain.EventTraceRecord{
		EventID:            values["event_id"],
		EventHash:          values["event_hash"],
		EventType:          values["event_type"],
		EventVersion:       values["event_version"],
		TruckID:            int32(truckID),
		StateCode:          values["state_code"],
		ResolvedAction:     values["resolved_action"],
		AdminNotified:      adminNotified,
		CitizenFanoutCount: citizenFanoutCount,
		Result:             values["result"],
		CreatedAt:          createdAt,
	}, nil
}
