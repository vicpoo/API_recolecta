package application

import (
	"github.com/vicpoo/API_recolecta/src/HistorialAsignacionCamion/domain/entities"
	"github.com/vicpoo/API_recolecta/src/HistorialAsignacionCamion/domain/ports"
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
