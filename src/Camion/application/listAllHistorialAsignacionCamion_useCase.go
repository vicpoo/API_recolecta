package application

import (
	"github.com/vicpoo/API_recolecta/src/Camion/domain/entities"
	"github.com/vicpoo/API_recolecta/src/Camion/domain/ports"
)

type ListAllHistorialAsignacionCamionUseCase struct {
	repo ports.IHistorialAsignacionCamion		
}

func NewListAllHistorialAsignacionCamionUseCase(repo ports.IHistorialAsignacionCamion) *ListAllHistorialAsignacionCamionUseCase {
	return &ListAllHistorialAsignacionCamionUseCase{
		repo: repo,
	}
}

func (uc *ListAllHistorialAsignacionCamionUseCase) Run() ([]entities.HistorialAsignacionCamion, error) {
	historial, err := uc.repo.ListAll()

	if err != nil {
		return nil, err
	}

	return historial, nil
}