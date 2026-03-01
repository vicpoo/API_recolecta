package application

import (
	"github.com/vicpoo/API_recolecta/src/Camion/domain/entities"
	"github.com/vicpoo/API_recolecta/src/Camion/domain/ports"
)

type GetActivoByCamionUseCase struct {
	repo ports.IHistorialAsignacionCamion
}

func NewGetActivoByCamionUseCase(repo ports.IHistorialAsignacionCamion) *GetActivoByCamionUseCase {
	return &GetActivoByCamionUseCase{repo: repo}
}

func (uc *GetActivoByCamionUseCase) Run(camionId int32) (*entities.HistorialAsignacionCamion, error) {
	return uc.repo.GetActivoByCamionId(camionId)
}
