package application

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"golang.org/x/crypto/bcrypt"
	"github.com/vicpoo/API_recolecta/src/ciudadano/domain/entities"
	"github.com/vicpoo/API_recolecta/src/ciudadano/domain/ports"
)

type RegisterCiudadanoRequest struct {
	Email     string   `json:"email"`
	Alias     string   `json:"alias"`
	Password  string   `json:"password"`
	FCMToken  string   `json:"fcm_token"`
	Latitude  *float64 `json:"latitude"`
	Longitude *float64 `json:"longitude"`
}

type RegisterCiudadanoUseCase struct {
	postgresRepo ports.CiudadanoPostgresRepository
	redisRepo    ports.CiudadanoRepository
}

func NewRegisterCiudadanoUseCase(
	postgresRepo ports.CiudadanoPostgresRepository,
	redisRepo ports.CiudadanoRepository,
) *RegisterCiudadanoUseCase {
	return &RegisterCiudadanoUseCase{postgresRepo: postgresRepo, redisRepo: redisRepo}
}

func (uc *RegisterCiudadanoUseCase) Execute(ctx context.Context, req RegisterCiudadanoRequest) (int, error) {
	req.Email = strings.TrimSpace(strings.ToLower(req.Email))

	existing, err := uc.postgresRepo.FindByEmail(ctx, req.Email)
	if err != nil {
		return 0, err
	}
	if existing != nil {
		return 0, errors.New("email ya registrado")
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return 0, err
	}

	ciudadano := &entities.CiudadanoPostgres{
		Email:        req.Email,
		Alias:        strings.TrimSpace(req.Alias),
		PasswordHash: string(hash),
	}

	id, err := uc.postgresRepo.Create(ctx, ciudadano)
	if err != nil {
		return 0, err
	}

	userIDStr := fmt.Sprintf("%d", id)

	if req.Latitude != nil && req.Longitude != nil {
		if err := uc.redisRepo.RegisterUser(ctx, userIDStr, *req.Longitude, *req.Latitude, req.FCMToken); err != nil {
			return id, err
		}
	} else {
		if err := uc.redisRepo.UpdateUserFCMToken(ctx, userIDStr, req.FCMToken); err != nil {
			return id, err
		}
	}

	return id, nil
}
