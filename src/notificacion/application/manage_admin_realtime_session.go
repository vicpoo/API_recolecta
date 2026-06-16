package application

import (
	"context"
	"time"

	"github.com/vicpoo/API_recolecta/src/notificacion/domain"
)

type ManageAdminRealtimeSessionUseCase struct {
	repo domain.IAdminRealtimeSessionRepository
}

func NewManageAdminRealtimeSessionUseCase(repo domain.IAdminRealtimeSessionRepository) *ManageAdminRealtimeSessionUseCase {
	return &ManageAdminRealtimeSessionUseCase{repo: repo}
}

func (uc *ManageAdminRealtimeSessionUseCase) IssueUpgradeToken(ctx context.Context, adminID int32, sessionID string) (*domain.AdminWSUpgradeTokenClaim, error) {
	epoch, err := uc.repo.GetOrCreateServerEpoch(ctx)
	if err != nil {
		return nil, err
	}
	now := time.Now().UTC()
	claim := &domain.AdminWSUpgradeTokenClaim{
		JTI:         sessionID,
		AdminID:     adminID,
		SessionID:   sessionID,
		ServerEpoch: epoch,
		IssuedAt:    now,
		ExpiresAt:   now.Add(5 * time.Minute),
	}
	if err := uc.repo.StoreUpgradeToken(ctx, claim, 5*time.Minute); err != nil {
		return nil, err
	}
	return claim, nil
}

func (uc *ManageAdminRealtimeSessionUseCase) ConsumeUpgradeToken(ctx context.Context, jti string) (*domain.AdminWSUpgradeTokenClaim, error) {
	return uc.repo.ConsumeUpgradeToken(ctx, jti)
}

func (uc *ManageAdminRealtimeSessionUseCase) RegisterSession(ctx context.Context, claim *domain.AdminWSUpgradeTokenClaim) (*domain.AdminWSSession, error) {
	now := time.Now().UTC()
	session := &domain.AdminWSSession{
		SessionID:   claim.SessionID,
		AdminID:     claim.AdminID,
		ServerEpoch: claim.ServerEpoch,
		LastSeenAt:  now,
		ConnectedAt: now,
		Status:      "active",
	}
	if err := uc.repo.UpsertSession(ctx, session, time.Hour); err != nil {
		return nil, err
	}
	return session, nil
}

func (uc *ManageAdminRealtimeSessionUseCase) Heartbeat(ctx context.Context, sessionID string) error {
	return uc.repo.TouchSession(ctx, sessionID, time.Now().UTC(), time.Hour)
}

func (uc *ManageAdminRealtimeSessionUseCase) GetSession(ctx context.Context, sessionID string) (*domain.AdminWSSession, error) {
	return uc.repo.GetSession(ctx, sessionID)
}

func (uc *ManageAdminRealtimeSessionUseCase) Disconnect(ctx context.Context, sessionID string) error {
	return uc.repo.InvalidateSession(ctx, sessionID)
}
