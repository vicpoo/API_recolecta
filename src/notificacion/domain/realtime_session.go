package domain

import (
	"context"
	"time"
)

type AdminWSSession struct {
	SessionID   string    `json:"session_id"`
	AdminID     int32     `json:"admin_id"`
	ServerEpoch string    `json:"server_epoch"`
	LastSeenAt  time.Time `json:"last_seen_at"`
	ConnectedAt time.Time `json:"connected_at"`
	Status      string    `json:"status"`
}

type AdminRealtimeSessionRepository interface {
	GetOrCreateServerEpoch(ctx context.Context) (string, error)
	StoreUpgradeToken(ctx context.Context, claim *AdminWSUpgradeTokenClaim, ttl time.Duration) error
	ConsumeUpgradeToken(ctx context.Context, jti string) (*AdminWSUpgradeTokenClaim, error)
	UpsertSession(ctx context.Context, session *AdminWSSession, ttl time.Duration) error
	TouchSession(ctx context.Context, sessionID string, lastSeenAt time.Time, ttl time.Duration) error
	GetSession(ctx context.Context, sessionID string) (*AdminWSSession, error)
	InvalidateSession(ctx context.Context, sessionID string) error
}
