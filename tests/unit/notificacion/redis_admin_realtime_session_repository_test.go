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

func setupRealtimeRepo(t *testing.T) (*notifinfra.RedisAdminRealtimeSessionRepository, context.Context) {
	t.Helper()
	mr := miniredis.RunT(t)
	rdb := goredis.NewClient(&goredis.Options{Addr: mr.Addr()})
	t.Cleanup(func() {
		_ = rdb.Close()
		mr.Close()
	})
	return notifinfra.NewRedisAdminRealtimeSessionRepository(rdb), context.Background()
}

func TestRedisAdminRealtimeSessionRepository_ServerEpochStable(t *testing.T) {
	repo, ctx := setupRealtimeRepo(t)
	first, err := repo.GetOrCreateServerEpoch(ctx)
	require.NoError(t, err)
	require.NotEmpty(t, first)

	second, err := repo.GetOrCreateServerEpoch(ctx)
	require.NoError(t, err)
	assert.Equal(t, first, second)
}

func TestRedisAdminRealtimeSessionRepository_ConsumeUpgradeTokenOneTime(t *testing.T) {
	repo, ctx := setupRealtimeRepo(t)
	now := time.Now().UTC()
	claim := &domain.AdminWSUpgradeTokenClaim{
		JTI:         "token-1",
		AdminID:     10,
		SessionID:   "session-1",
		ServerEpoch: "epoch-1",
		IssuedAt:    now,
		ExpiresAt:   now.Add(5 * time.Minute),
	}

	require.NoError(t, repo.StoreUpgradeToken(ctx, claim, 5*time.Minute))

	consumed, err := repo.ConsumeUpgradeToken(ctx, "token-1")
	require.NoError(t, err)
	assert.Equal(t, int32(10), consumed.AdminID)
	assert.Equal(t, "session-1", consumed.SessionID)

	_, err = repo.ConsumeUpgradeToken(ctx, "token-1")
	require.Error(t, err)
}

func TestRedisAdminRealtimeSessionRepository_SessionLifecycle(t *testing.T) {
	repo, ctx := setupRealtimeRepo(t)
	now := time.Now().UTC().Truncate(time.Second)
	session := &domain.AdminWSSession{
		SessionID:   "session-2",
		AdminID:     22,
		ServerEpoch: "epoch-2",
		LastSeenAt:  now,
		ConnectedAt: now,
		Status:      "active",
	}

	require.NoError(t, repo.UpsertSession(ctx, session, time.Hour))

	stored, err := repo.GetSession(ctx, "session-2")
	require.NoError(t, err)
	assert.Equal(t, int32(22), stored.AdminID)
	assert.Equal(t, "active", stored.Status)

	require.NoError(t, repo.TouchSession(ctx, "session-2", now.Add(2*time.Minute), time.Hour))

	require.NoError(t, repo.InvalidateSession(ctx, "session-2"))
	_, err = repo.GetSession(ctx, "session-2")
	require.Error(t, err)
}
