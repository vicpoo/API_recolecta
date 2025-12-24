package application

import (
	"github.com/vicpoo/API_recolecta/src/Camion/domain/entities"
	"github.com/vicpoo/API_recolecta/src/Camion/domain/ports"
)

type ListCamionUseCase struct {
	repo ports.ICamion
}

func NewListCamionUseCase(repo ports.ICamion) *ListCamionUseCase {
	return &ListCamionUseCase{
		repo: repo,
	}
}

func (uc *ListCamionUseCase) Run() ([]entities.Camion, error) {
	return uc.repo.ListAll()
}
