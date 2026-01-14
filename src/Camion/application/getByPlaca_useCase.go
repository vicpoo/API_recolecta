package application

import (
	"errors"

	"github.com/vicpoo/API_recolecta/src/Camion/domain/entities"
	"github.com/vicpoo/API_recolecta/src/Camion/domain/ports"
)

type GetCamionByPlacaUseCase struct {
	repo ports.ICamion
}

func NewGetCamionByPlacaUseCase(repo ports.ICamion) *GetCamionByPlacaUseCase {
	return &GetCamionByPlacaUseCase{repo: repo}
}

func (uc *GetCamionByPlacaUseCase) Run(placa string) (*entities.Camion, error) {
	if placa == "" {
		return nil, errors.New("la placa es obligatoria")
	}
	return uc.repo.GetByPlaca(placa)
}
