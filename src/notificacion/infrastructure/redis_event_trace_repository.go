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
