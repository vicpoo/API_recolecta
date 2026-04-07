package application

import (
	"context"
	"errors"
	"fmt"

	"github.com/vicpoo/API_recolecta/src/ciudadano/domain/ports"
)

type UpdateCoordinatesRequest struct {
	UserID    int
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type UpdateCoordinatesUseCase struct {
	redisRepo    ports.CiudadanoRepository
	postgresRepo ports.CiudadanoPostgresRepository
}

func NewUpdateCoordinatesUseCase(
	redisRepo ports.CiudadanoRepository,
	postgresRepo ports.CiudadanoPostgresRepository,
) *UpdateCoordinatesUseCase {
	return &UpdateCoordinatesUseCase{redisRepo: redisRepo, postgresRepo: postgresRepo}
}

func (uc *UpdateCoordinatesUseCase) Execute(ctx context.Context, req UpdateCoordinatesRequest) error {
	citizen, err := uc.postgresRepo.FindByID(ctx, req.UserID)
	if err != nil {
		return err
	}
	if citizen == nil {
		return errors.New("ciudadano no encontrado")
	}

	userIDStr := fmt.Sprintf("%d", req.UserID)
	return uc.redisRepo.UpdateUserGeoCoordinates(ctx, userIDStr, req.Longitude, req.Latitude)
}
