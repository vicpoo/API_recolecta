package infrastructure

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
	"github.com/vicpoo/API_recolecta/src/notificacion/domain"
)

type RedisAdminRealtimeSessionRepository struct {
	rdb *redis.Client
}

func NewRedisAdminRealtimeSessionRepository(rdb *redis.Client) *RedisAdminRealtimeSessionRepository {
	return &RedisAdminRealtimeSessionRepository{rdb: rdb}
}

func (r *RedisAdminRealtimeSessionRepository) GetOrCreateServerEpoch(ctx context.Context) (string, error) {
	const key = "ws:server:epoch"
	epoch := uuid.New().String()
	ok, err := r.rdb.SetNX(ctx, key, epoch, 0).Result()
	if err != nil {
		return "", err
	}
	if ok {
		return epoch, nil
	}
	return r.rdb.Get(ctx, key).Result()
}

func (r *RedisAdminRealtimeSessionRepository) StoreUpgradeToken(ctx context.Context, claim *domain.AdminWSUpgradeTokenClaim, ttl time.Duration) error {
	data, err := json.Marshal(claim)
	if err != nil {
		return err
	}
	key := fmt.Sprintf("ws:upgrade:%s", claim.JTI)
	return r.rdb.Set(ctx, key, data, ttl).Err()
}

func (r *RedisAdminRealtimeSessionRepository) ConsumeUpgradeToken(ctx context.Context, jti string) (*domain.AdminWSUpgradeTokenClaim, error) {
	key := fmt.Sprintf("ws:upgrade:%s", jti)
	val, err := r.rdb.GetDel(ctx, key).Result()
	if err == redis.Nil {
		return nil, fmt.Errorf("token not found or already consumed: %s", jti)
	}
	if err != nil {
		return nil, err
	}
	var claim domain.AdminWSUpgradeTokenClaim
	if err := json.Unmarshal([]byte(val), &claim); err != nil {
		return nil, err
	}
	return &claim, nil
}

func (r *RedisAdminRealtimeSessionRepository) UpsertSession(ctx context.Context, session *domain.AdminWSSession, ttl time.Duration) error {
	data, err := json.Marshal(session)
	if err != nil {
		return err
	}
	key := fmt.Sprintf("ws:session:%s", session.SessionID)
	return r.rdb.Set(ctx, key, data, ttl).Err()
}

func (r *RedisAdminRealtimeSessionRepository) GetSession(ctx context.Context, sessionID string) (*domain.AdminWSSession, error) {
	key := fmt.Sprintf("ws:session:%s", sessionID)
	val, err := r.rdb.Get(ctx, key).Result()
	if err == redis.Nil {
		return nil, fmt.Errorf("session not found: %s", sessionID)
	}
	if err != nil {
		return nil, err
	}
	var session domain.AdminWSSession
	if err := json.Unmarshal([]byte(val), &session); err != nil {
		return nil, err
	}
	return &session, nil
}

func (r *RedisAdminRealtimeSessionRepository) TouchSession(ctx context.Context, sessionID string, lastSeenAt time.Time, ttl time.Duration) error {
	session, err := r.GetSession(ctx, sessionID)
	if err != nil {
		return err
	}
	session.LastSeenAt = lastSeenAt
	return r.UpsertSession(ctx, session, ttl)
}

func (r *RedisAdminRealtimeSessionRepository) InvalidateSession(ctx context.Context, sessionID string) error {
	key := fmt.Sprintf("ws:session:%s", sessionID)
	return r.rdb.Del(ctx, key).Err()
}
