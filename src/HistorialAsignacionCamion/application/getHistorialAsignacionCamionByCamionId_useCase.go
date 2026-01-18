package application

import (
	"github.com/vicpoo/API_recolecta/src/HistorialAsignacionCamion/domain/entities"
	"github.com/vicpoo/API_recolecta/src/HistorialAsignacionCamion/domain/ports"
)

type GetHistorialByCamionUseCase struct {
	repo ports.IHistorialAsignacionCamion
}

func NewGetHistorialByCamionUseCase(repo ports.IHistorialAsignacionCamion) *GetHistorialByCamionUseCase {
	return &GetHistorialByCamionUseCase{repo: repo}
}

func (uc *GetHistorialByCamionUseCase) Run(camionId int32) ([]entities.HistorialAsignacionCamion, error) {
	return uc.repo.GetByCamionId(camionId)
}
