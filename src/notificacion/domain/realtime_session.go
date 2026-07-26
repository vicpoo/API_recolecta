package domain

import (
	"context"
	"time"
)

type AdminWSUpgradeTokenClaim struct {
	JTI         string
	AdminID     int32
	SessionID   string
	ServerEpoch string
	IssuedAt    time.Time
	ExpiresAt   time.Time
}

type AdminWSSession struct {
	SessionID   string
	AdminID     int32
	ServerEpoch string
	LastSeenAt  time.Time
	ConnectedAt time.Time
	Status      string
}

type IAdminRealtimeSessionRepository interface {
	GetOrCreateServerEpoch(ctx context.Context) (string, error)
	StoreUpgradeToken(ctx context.Context, claim *AdminWSUpgradeTokenClaim, ttl time.Duration) error
	ConsumeUpgradeToken(ctx context.Context, jti string) (*AdminWSUpgradeTokenClaim, error)
	UpsertSession(ctx context.Context, session *AdminWSSession, ttl time.Duration) error
	GetSession(ctx context.Context, sessionID string) (*AdminWSSession, error)
	TouchSession(ctx context.Context, sessionID string, lastSeenAt time.Time, ttl time.Duration) error
	InvalidateSession(ctx context.Context, sessionID string) error
}
