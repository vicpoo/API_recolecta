package application

import (
	"github.com/vicpoo/API_recolecta/src/HistorialAsignacionCamion/domain/entities"
	"github.com/vicpoo/API_recolecta/src/HistorialAsignacionCamion/domain/ports"
)

type SaveHistorialAsignacionCamionUseCase struct {
	repo ports.IHistorialAsignacionCamion
}

func NewSaveHistorialAsignacionCamionUseCase(repo ports.IHistorialAsignacionCamion) *SaveHistorialAsignacionCamionUseCase {
	return &SaveHistorialAsignacionCamionUseCase{
		repo: repo,
	}
}

func (uc *SaveHistorialAsignacionCamionUseCase) Run(historial *entities.HistorialAsignacionCamion) (*entities.HistorialAsignacionCamion, error){
	historial, err := uc.repo.Save(historial)

	if err != nil {
		return nil, err
	}

	return historial, nil
}