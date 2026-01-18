package application

import (
	"github.com/vicpoo/API_recolecta/src/RutaCamion/domain/entities"
	"github.com/vicpoo/API_recolecta/src/RutaCamion/domain/ports"
)

type ListAllRutaCamionUseCase struct {
	repo ports.RutaCamionRepository
}

func NewListAllRutaCamionUseCase(repo ports.RutaCamionRepository) *ListAllRutaCamionUseCase {
	return &ListAllRutaCamionUseCase{repo: repo}
}

func (uc *ListAllRutaCamionUseCase) Execute() ([]entities.RutaCamion, error) {
	return uc.repo.ListAll()
}
