package application

import (
	"github.com/vicpoo/API_recolecta/src/HistorialAsignacionCamion/domain/entities"
	"github.com/vicpoo/API_recolecta/src/HistorialAsignacionCamion/domain/ports"
)

type GetHistorialByChoferUseCase struct {
	repo ports.IHistorialAsignacionCamion
}

func NewGetHistorialByChoferUseCase(repo ports.IHistorialAsignacionCamion) *GetHistorialByChoferUseCase {
	return &GetHistorialByChoferUseCase{repo: repo}
}

func (uc *GetHistorialByChoferUseCase) Run(choferId int32) ([]entities.HistorialAsignacionCamion, error) {
	return uc.repo.GetByChoferId(choferId)
}
