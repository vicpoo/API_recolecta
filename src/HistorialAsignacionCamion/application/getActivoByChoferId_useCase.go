package application

import (
	"github.com/vicpoo/API_recolecta/src/HistorialAsignacionCamion/domain/entities"
	"github.com/vicpoo/API_recolecta/src/HistorialAsignacionCamion/domain/ports"
)

type GetActivoByChoferUseCase struct {
	repo ports.IHistorialAsignacionCamion
}

func NewGetActivoByChoferUseCase(repo ports.IHistorialAsignacionCamion) *GetActivoByChoferUseCase {
	return &GetActivoByChoferUseCase{repo: repo}
}

func (uc *GetActivoByChoferUseCase) Run(choferId int32) (*entities.HistorialAsignacionCamion, error) {
	return uc.repo.GetActivoByChoferId(choferId)
}
