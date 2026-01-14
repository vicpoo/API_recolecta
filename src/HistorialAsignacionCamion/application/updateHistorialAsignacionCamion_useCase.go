package application

import (
	"github.com/vicpoo/API_recolecta/src/HistorialAsignacionCamion/domain/entities"
	"github.com/vicpoo/API_recolecta/src/HistorialAsignacionCamion/domain/ports"
)

type UpdateHistorialAsignacionCamionUseCase struct {
	repo ports.IHistorialAsignacionCamion
}

func NewUpdateHistorialAsignacionCamionUseCase(repo ports.IHistorialAsignacionCamion) *UpdateHistorialAsignacionCamionUseCase {
	return &UpdateHistorialAsignacionCamionUseCase{
		repo: repo,
	}
}

func (uc *UpdateHistorialAsignacionCamionUseCase) Run(id int32, historial *entities.HistorialAsignacionCamion) (*entities.HistorialAsignacionCamion, error) {
	historialCamion, err := uc.repo.Update(id, historial)

	if err != nil {
		return nil, err
	}

	return historialCamion, nil
}
