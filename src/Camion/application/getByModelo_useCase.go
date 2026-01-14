package application

import (
	"errors"

	"github.com/vicpoo/API_recolecta/src/Camion/domain/entities"
	"github.com/vicpoo/API_recolecta/src/Camion/domain/ports"
)

type GetCamionByModeloUseCase struct {
	repo ports.ICamion
}

func NewGetCamionByModeloUseCase(repo ports.ICamion) *GetCamionByModeloUseCase {
	return &GetCamionByModeloUseCase{repo: repo}
}

func (uc *GetCamionByModeloUseCase) Run(modelo string) ([]entities.Camion, error) {
	if modelo == "" {
		return nil, errors.New("el modelo es obligatorio")
	}
	return uc.repo.GetByModelo(modelo)
}
