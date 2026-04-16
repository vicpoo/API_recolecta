package notificacion_test

import (
	"context"
	"testing"
	"time"

	"github.com/alicebob/miniredis/v2"
	goredis "github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/vicpoo/API_recolecta/src/notificacion/domain"
	notifinfra "github.com/vicpoo/API_recolecta/src/notificacion/infrastructure"
)

func setupEventTraceRepo(t *testing.T) (*notifinfra.RedisEventTraceRepository, context.Context) {
	t.Helper()
	mr := miniredis.RunT(t)
	rdb := goredis.NewClient(&goredis.Options{Addr: mr.Addr()})
	t.Cleanup(func() {
		_ = rdb.Close()
		mr.Close()
	})
	return notifinfra.NewRedisEventTraceRepository(rdb), context.Background()
}

func TestRedisEventTraceRepository_TryAcquireDeduplication(t *testing.T) {
	repo, ctx := setupEventTraceRepo(t)
	event := &domain.TruckStateEvent{
		EventID:      "evt-1",
		EventType:    "TRUCK_STATE_CHANGED",
		EventVersion: domain.EventVersionV1,
		TruckID:      5,
		OccurredAt:   time.Now().UTC(),
	}

	first, err := repo.TryAcquireDeduplication(ctx, "hash-abc", event)
	require.NoError(t, err)
	assert.True(t, first)

	second, err := repo.TryAcquireDeduplication(ctx, "hash-abc", event)
	require.NoError(t, err)
	assert.False(t, second)
}

func TestRedisEventTraceRepository_SaveAndQuery(t *testing.T) {
	repo, ctx := setupEventTraceRepo(t)
	trace := &domain.EventTraceRecord{
		EventID:            "evt-2",
		EventHash:          "hash-2",
		EventType:          "TRUCK_STATE_CHANGED",
		EventVersion:       domain.EventVersionV1,
		TruckID:            8,
		StateCode:          domain.StateArrival,
		ResolvedAction:     domain.ActionNotifyAdminOnly,
		AdminNotified:      true,
		CitizenFanoutCount: 0,
		Result:             "processed",
		CreatedAt:          time.Now().UTC().Truncate(time.Second),
	}

	require.NoError(t, repo.SaveTrace(ctx, trace))

	stored, err := repo.GetByEventID(ctx, "evt-2")
	require.NoError(t, err)
	assert.Equal(t, trace.EventID, stored.EventID)
	assert.Equal(t, trace.TruckID, stored.TruckID)
	assert.Equal(t, trace.StateCode, stored.StateCode)

	list, err := repo.ListByTruckID(ctx, 8, 10)
	require.NoError(t, err)
	require.Len(t, list, 1)
	assert.Equal(t, "evt-2", list[0].EventID)
}
