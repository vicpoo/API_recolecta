package application

import (
	"github.com/vicpoo/API_recolecta/src/RutaCamion/domain/entities"
	"github.com/vicpoo/API_recolecta/src/RutaCamion/domain/ports"
)

type GetRutaCamionByRutaIDUseCase struct {
	repo ports.RutaCamionRepository
}

func NewGetRutaCamionByRutaIDUseCase(repo ports.RutaCamionRepository) *GetRutaCamionByRutaIDUseCase {
	return &GetRutaCamionByRutaIDUseCase{repo: repo}
}

func (uc *GetRutaCamionByRutaIDUseCase) Execute(
	rutaID int32,
) ([]entities.RutaCamion, error) {
	return uc.repo.GetByRutaID(rutaID)
}
