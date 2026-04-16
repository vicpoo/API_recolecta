package application

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/vicpoo/API_recolecta/src/notificacion/domain"
)

const (
	wsUpgradeTokenTTL = 5 * time.Minute
	wsSessionTTL      = 1 * time.Hour
)

type IssueUpgradeTokenOutput struct {
	WSUpgradeToken string    `json:"ws_upgrade_token"`
	SessionID      string    `json:"session_id"`
	ServerEpoch    string    `json:"server_epoch"`
	ExpiresAt      time.Time `json:"expires_at"`
}

type ConsumeUpgradeTokenOutput struct {
	SessionID   string `json:"session_id"`
	ServerEpoch string `json:"server_epoch"`
	Status      string `json:"status"`
}

type ManageAdminRealtimeSessionUseCase struct {
	repo domain.AdminRealtimeSessionRepository
}

func NewManageAdminRealtimeSessionUseCase(repo domain.AdminRealtimeSessionRepository) *ManageAdminRealtimeSessionUseCase {
	return &ManageAdminRealtimeSessionUseCase{repo: repo}
}

func (uc *ManageAdminRealtimeSessionUseCase) IssueUpgradeToken(ctx context.Context, adminID int32) (*IssueUpgradeTokenOutput, error) {
	if adminID <= 0 {
		return nil, fmt.Errorf("admin_id invalido")
	}

	serverEpoch, err := uc.repo.GetOrCreateServerEpoch(ctx)
	if err != nil {
		return nil, err
	}

	now := time.Now().UTC()
	claim := &domain.AdminWSUpgradeTokenClaim{
		JTI:         uuid.NewString(),
		AdminID:     adminID,
		SessionID:   uuid.NewString(),
		ServerEpoch: serverEpoch,
		IssuedAt:    now,
		ExpiresAt:   now.Add(wsUpgradeTokenTTL),
	}

	if err := uc.repo.StoreUpgradeToken(ctx, claim, wsUpgradeTokenTTL); err != nil {
		return nil, err
	}

	return &IssueUpgradeTokenOutput{
		WSUpgradeToken: claim.JTI,
		SessionID:      claim.SessionID,
		ServerEpoch:    claim.ServerEpoch,
		ExpiresAt:      claim.ExpiresAt,
	}, nil
}

func (uc *ManageAdminRealtimeSessionUseCase) ConsumeUpgradeToken(ctx context.Context, wsUpgradeToken string) (*ConsumeUpgradeTokenOutput, error) {
	token := strings.TrimSpace(wsUpgradeToken)
	if token == "" {
		return nil, fmt.Errorf("ws_upgrade_token es obligatorio")
	}

	currentEpoch, err := uc.repo.GetOrCreateServerEpoch(ctx)
	if err != nil {
		return nil, err
	}

	claim, err := uc.repo.ConsumeUpgradeToken(ctx, token)
	if err != nil {
		return nil, err
	}

	if claim.ServerEpoch != currentEpoch {
		return nil, fmt.Errorf("token invalido por epoch de servidor")
	}

	now := time.Now().UTC()
	session := &domain.AdminWSSession{
		SessionID:   claim.SessionID,
		AdminID:     claim.AdminID,
		ServerEpoch: claim.ServerEpoch,
		LastSeenAt:  now,
		ConnectedAt: now,
		Status:      "active",
	}
	if err := uc.repo.UpsertSession(ctx, session, wsSessionTTL); err != nil {
		return nil, err
	}

	return &ConsumeUpgradeTokenOutput{
		SessionID:   session.SessionID,
		ServerEpoch: session.ServerEpoch,
		Status:      session.Status,
	}, nil
}

func (uc *ManageAdminRealtimeSessionUseCase) Heartbeat(ctx context.Context, sessionID string) error {
	normalized := strings.TrimSpace(sessionID)
	if normalized == "" {
		return fmt.Errorf("session_id es obligatorio")
	}
	return uc.repo.TouchSession(ctx, normalized, time.Now().UTC(), wsSessionTTL)
}

func (uc *ManageAdminRealtimeSessionUseCase) Disconnect(ctx context.Context, sessionID string) error {
	normalized := strings.TrimSpace(sessionID)
	if normalized == "" {
		return fmt.Errorf("session_id es obligatorio")
	}
	return uc.repo.InvalidateSession(ctx, normalized)
}

func (uc *ManageAdminRealtimeSessionUseCase) GetSession(ctx context.Context, sessionID string) (*domain.AdminWSSession, error) {
	normalized := strings.TrimSpace(sessionID)
	if normalized == "" {
		return nil, fmt.Errorf("session_id es obligatorio")
	}
	return uc.repo.GetSession(ctx, normalized)
}
