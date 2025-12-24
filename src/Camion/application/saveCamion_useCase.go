package application

import (
	"errors"

	"github.com/vicpoo/API_recolecta/src/Camion/domain/entities"
	"github.com/vicpoo/API_recolecta/src/Camion/domain/ports"
)

type SaveCamionUseCase struct {
	repo ports.ICamion
}

func NewSaveCamionUseCase(repo ports.ICamion) *SaveCamionUseCase {
	return &SaveCamionUseCase{
		repo: repo,
	}
}

func (uc *SaveCamionUseCase) Run(c *entities.Camion) (*entities.Camion, error) {
	if c.Placa == "" || c.Modelo == "" {
		return nil, errors.New("placa y modelo son obligatorios")
	}
	return uc.repo.Save(c)
}
