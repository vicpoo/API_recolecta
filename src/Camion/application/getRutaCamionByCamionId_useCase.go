package application

import (
	"github.com/vicpoo/API_recolecta/src/Camion/domain/entities"
	"github.com/vicpoo/API_recolecta/src/Camion/domain/ports"
)

type GetRutaCamionByCamionIDUseCase struct {
	repo ports.RutaCamionRepository
}

func NewGetRutaCamionByCamionIDUseCase(repo ports.RutaCamionRepository) *GetRutaCamionByCamionIDUseCase {
	return &GetRutaCamionByCamionIDUseCase{repo: repo}
}

func (uc *GetRutaCamionByCamionIDUseCase) Execute(
	camionID int32,
) ([]entities.RutaCamion, error) {
	return uc.repo.GetByCamionID(camionID)
}
