package infrastructure

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/google/uuid"
	goredis "github.com/redis/go-redis/v9"
	"github.com/vicpoo/API_recolecta/src/notificacion/domain"
)

const (
	realtimeServerEpochKey = "realtime:server_epoch:current"
	wsUpgradeKeyFmt        = "ws:upgrade:%s"
	wsSessionKeyFmt        = "ws:session:%s"
)

type RedisAdminRealtimeSessionRepository struct {
	rdb *goredis.Client
}

func NewRedisAdminRealtimeSessionRepository(rdb *goredis.Client) *RedisAdminRealtimeSessionRepository {
	return &RedisAdminRealtimeSessionRepository{rdb: rdb}
}

func (r *RedisAdminRealtimeSessionRepository) GetOrCreateServerEpoch(ctx context.Context) (string, error) {
	epoch, err := r.rdb.Get(ctx, realtimeServerEpochKey).Result()
	if err == nil {
		return epoch, nil
	}
	if err != goredis.Nil {
		return "", fmt.Errorf("no se pudo leer server epoch: %w", err)
	}

	newEpoch := uuid.NewString()
	if err := r.rdb.SetNX(ctx, realtimeServerEpochKey, newEpoch, 0).Err(); err != nil {
		return "", fmt.Errorf("no se pudo crear server epoch: %w", err)
	}

	epoch, err = r.rdb.Get(ctx, realtimeServerEpochKey).Result()
	if err != nil {
		return "", fmt.Errorf("no se pudo confirmar server epoch: %w", err)
	}
	return epoch, nil
}

func (r *RedisAdminRealtimeSessionRepository) StoreUpgradeToken(ctx context.Context, claim *domain.AdminWSUpgradeTokenClaim, ttl time.Duration) error {
	key := fmt.Sprintf(wsUpgradeKeyFmt, claim.JTI)
	fields := map[string]interface{}{
		"admin_id":     claim.AdminID,
		"session_id":   claim.SessionID,
		"server_epoch": claim.ServerEpoch,
		"issued_at":    claim.IssuedAt.UTC().Format(time.RFC3339),
		"expires_at":   claim.ExpiresAt.UTC().Format(time.RFC3339),
		"used":         "false",
	}

	_, err := r.rdb.TxPipelined(ctx, func(pipe goredis.Pipeliner) error {
		pipe.HSet(ctx, key, fields)
		pipe.Expire(ctx, key, ttl)
		return nil
	})
	if err != nil {
		return fmt.Errorf("no se pudo guardar ws upgrade token: %w", err)
	}
	return nil
}

func (r *RedisAdminRealtimeSessionRepository) ConsumeUpgradeToken(ctx context.Context, jti string) (*domain.AdminWSUpgradeTokenClaim, error) {
	key := fmt.Sprintf(wsUpgradeKeyFmt, jti)

	var claim *domain.AdminWSUpgradeTokenClaim
	err := r.rdb.Watch(ctx, func(tx *goredis.Tx) error {
		values, getErr := tx.HGetAll(ctx, key).Result()
		if getErr != nil {
			return fmt.Errorf("no se pudo leer upgrade token: %w", getErr)
		}
		if len(values) == 0 {
			return fmt.Errorf("upgrade token no encontrado")
		}
		if values["used"] == "true" {
			return fmt.Errorf("upgrade token ya utilizado")
		}

		adminID, convErr := strconv.Atoi(values["admin_id"])
		if convErr != nil {
			return fmt.Errorf("admin_id invalido en upgrade token: %w", convErr)
		}
		issuedAt, convErr := time.Parse(time.RFC3339, values["issued_at"])
		if convErr != nil {
			return fmt.Errorf("issued_at invalido en upgrade token: %w", convErr)
		}
		expiresAt, convErr := time.Parse(time.RFC3339, values["expires_at"])
		if convErr != nil {
			return fmt.Errorf("expires_at invalido en upgrade token: %w", convErr)
		}
		if time.Now().UTC().After(expiresAt) {
			return fmt.Errorf("upgrade token expirado")
		}

		claim = &domain.AdminWSUpgradeTokenClaim{
			JTI:         jti,
			AdminID:     int32(adminID),
			SessionID:   values["session_id"],
			ServerEpoch: values["server_epoch"],
			IssuedAt:    issuedAt,
			ExpiresAt:   expiresAt,
		}

		_, pipeErr := tx.TxPipelined(ctx, func(pipe goredis.Pipeliner) error {
			pipe.HSet(ctx, key, "used", "true")
			return nil
		})
		return pipeErr
	}, key)
	if err != nil {
		return nil, err
	}

	return claim, nil
}

func (r *RedisAdminRealtimeSessionRepository) UpsertSession(ctx context.Context, session *domain.AdminWSSession, ttl time.Duration) error {
	key := fmt.Sprintf(wsSessionKeyFmt, session.SessionID)
	fields := map[string]interface{}{
		"admin_id":     session.AdminID,
		"server_epoch": session.ServerEpoch,
		"last_seen_at": session.LastSeenAt.UTC().Format(time.RFC3339),
		"connected_at": session.ConnectedAt.UTC().Format(time.RFC3339),
		"status":       session.Status,
	}

	_, err := r.rdb.TxPipelined(ctx, func(pipe goredis.Pipeliner) error {
		pipe.HSet(ctx, key, fields)
		pipe.Expire(ctx, key, ttl)
		return nil
	})
	if err != nil {
		return fmt.Errorf("no se pudo guardar ws session: %w", err)
	}
	return nil
}

func (r *RedisAdminRealtimeSessionRepository) TouchSession(ctx context.Context, sessionID string, lastSeenAt time.Time, ttl time.Duration) error {
	key := fmt.Sprintf(wsSessionKeyFmt, sessionID)
	exists, err := r.rdb.Exists(ctx, key).Result()
	if err != nil {
		return fmt.Errorf("no se pudo validar session: %w", err)
	}
	if exists == 0 {
		return fmt.Errorf("session no encontrada")
	}

	_, err = r.rdb.TxPipelined(ctx, func(pipe goredis.Pipeliner) error {
		pipe.HSet(ctx, key,
			"last_seen_at", lastSeenAt.UTC().Format(time.RFC3339),
			"status", "active",
		)
		pipe.Expire(ctx, key, ttl)
		return nil
	})
	if err != nil {
		return fmt.Errorf("no se pudo registrar heartbeat: %w", err)
	}
	return nil
}

func (r *RedisAdminRealtimeSessionRepository) InvalidateSession(ctx context.Context, sessionID string) error {
	key := fmt.Sprintf(wsSessionKeyFmt, sessionID)
	if err := r.rdb.HSet(ctx, key, "status", "closed").Err(); err != nil {
		return fmt.Errorf("no se pudo cerrar session: %w", err)
	}
	if err := r.rdb.Del(ctx, key).Err(); err != nil {
		return fmt.Errorf("no se pudo invalidar session: %w", err)
	}
	return nil
}
