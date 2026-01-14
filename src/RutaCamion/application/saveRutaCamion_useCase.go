package application

import (
	"github.com/vicpoo/API_recolecta/src/RutaCamion/domain/entities"
	"github.com/vicpoo/API_recolecta/src/RutaCamion/domain/ports"
)

type SaveRutaCamionUseCase struct {
	repo ports.RutaCamionRepository
}

func NewSaveRutaCamionUseCase(repo ports.RutaCamionRepository) *SaveRutaCamionUseCase {
	return &SaveRutaCamionUseCase{repo: repo}
}

func (uc *SaveRutaCamionUseCase) Execute(
	rutaCamion *entities.RutaCamion,
) (*entities.RutaCamion, error) {
	return uc.repo.Save(rutaCamion)
}
