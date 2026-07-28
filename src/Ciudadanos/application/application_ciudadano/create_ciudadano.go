package application_ciudadano

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/vicpoo/API_recolecta/src/Ciudadanos/domain"
	"github.com/vicpoo/API_recolecta/src/Ciudadanos/domain/entities"
	"github.com/vicpoo/API_recolecta/src/core"
	passwordSecurity "github.com/vicpoo/API_recolecta/src/security/password"
)

type CreateCiudadanoInput struct {
	Email    string `json:"email"`
	Alias    string `json:"alias"`
	Password string `json:"password"`
	FCMToken string `json:"fcm_token"`
}

type CreateCiudadano struct {
	repo domain.CiudadanoRepository
}

func NewCreateCiudadano(repo domain.CiudadanoRepository) *CreateCiudadano {
	return &CreateCiudadano{repo: repo}
}

func (uc *CreateCiudadano) Execute(ctx context.Context, in CreateCiudadanoInput) (int, error) {
	in.Email = strings.TrimSpace(strings.ToLower(in.Email))
	in.Alias = strings.TrimSpace(in.Alias)
	in.Password = strings.TrimSpace(in.Password)
	in.FCMToken = strings.TrimSpace(in.FCMToken)

	if in.Email == "" {
		return 0, errors.New("email es requerido")
	}
	if in.Alias == "" {
		return 0, errors.New("alias es requerido")
	}
	if in.Password == "" {
		return 0, errors.New("password es requerido")
	}
	if in.FCMToken == "" {
		return 0, errors.New("fcm_token es requerido")
	}

	existingByEmail, err := uc.repo.FindByEmail(ctx, in.Email)
	if err != nil {
		return 0, err
	}
	if existingByEmail != nil {
		return 0, errors.New("el email ya está registrado")
	}

	existingByAlias, err := uc.repo.FindByAlias(ctx, in.Alias)
	if err != nil {
		return 0, err
	}
	if existingByAlias != nil {
		return 0, errors.New("el alias ya está registrado")
	}

	hash, err := passwordSecurity.Hash(in.Password)
	if err != nil {
		return 0, err
	}

	ciudadano := &entities.Ciudadano{
		Email:     in.Email,
		Alias:     in.Alias,
		Password:  hash,
		CreatedAt: time.Now(),
	}

	id, err := uc.repo.Create(ctx, ciudadano)
	if err != nil {
		return 0, err
	}

	rdb, err := core.ConnectRedis()
	if err != nil {
		_ = uc.repo.Delete(ctx, id)
		return 0, err
	}

	legacyKey := fmt.Sprintf("fcm:ciudadano:%d", id)
	userKey := fmt.Sprintf("user:%d", id)

	pipe := rdb.TxPipeline()
	pipe.Set(ctx, legacyKey, in.FCMToken, 0)
	pipe.HSet(ctx, userKey,
		"fcm_token", in.FCMToken,
		"fcm_status", "valid",
		"updated_at", time.Now().UTC().Format(time.RFC3339),
	)

	if _, err := pipe.Exec(ctx); err != nil {
		_ = uc.repo.Delete(ctx, id)
		return 0, err
	}

	return id, nil
}
