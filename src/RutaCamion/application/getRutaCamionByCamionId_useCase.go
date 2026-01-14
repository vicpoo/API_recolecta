package application

import (
	"github.com/vicpoo/API_recolecta/src/RutaCamion/domain/entities"
	"github.com/vicpoo/API_recolecta/src/RutaCamion/domain/ports"
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
