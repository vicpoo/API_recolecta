package application

import (
	"github.com/vicpoo/API_recolecta/src/HistorialAsignacionCamion/domain/entities"
	"github.com/vicpoo/API_recolecta/src/HistorialAsignacionCamion/domain/ports"
)

type GetHistorialAsignacionCamionByIdUseCase struct {
	repo ports.IHistorialAsignacionCamion
}

func NewGetHistorialAsignacionCamionByIdUseCase(repo ports.IHistorialAsignacionCamion) *GetHistorialAsignacionCamionByIdUseCase {
	return &GetHistorialAsignacionCamionByIdUseCase{repo: repo}
}

func (uc *GetHistorialAsignacionCamionByIdUseCase) Run(id int32) (*entities.HistorialAsignacionCamion, error) {
	return uc.repo.GetById(id)
}
