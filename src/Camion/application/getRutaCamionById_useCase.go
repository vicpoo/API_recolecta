package application

import (
	"github.com/vicpoo/API_recolecta/src/Camion/domain/entities"
	"github.com/vicpoo/API_recolecta/src/Camion/domain/ports"
)

type GetRutaCamionByIDUseCase struct {
	repo ports.RutaCamionRepository
}

func NewGetRutaCamionByIDUseCase(
	repo ports.RutaCamionRepository,
) *GetRutaCamionByIDUseCase {
	return &GetRutaCamionByIDUseCase{repo: repo}
}

func (uc *GetRutaCamionByIDUseCase) Execute(
	id int32,
) (*entities.RutaCamion, error) {
	return uc.repo.GetByID(id)
}
