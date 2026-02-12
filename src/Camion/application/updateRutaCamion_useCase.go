package application

import (
	"github.com/vicpoo/API_recolecta/src/Camion/domain/entities"
	"github.com/vicpoo/API_recolecta/src/Camion/domain/ports"
)

type UpdateRutaCamionUseCase struct {
	repo ports.RutaCamionRepository
}

func NewUpdateRutaCamionUseCase(repo ports.RutaCamionRepository) *UpdateRutaCamionUseCase {
	return &UpdateRutaCamionUseCase{repo: repo}
}

func (uc *UpdateRutaCamionUseCase) Execute(
	id int32,
	rutaCamion *entities.RutaCamion,
) (*entities.RutaCamion, error) {
	return uc.repo.Update(id, rutaCamion)
}
