package application

import (
	"errors"

	"github.com/vicpoo/API_recolecta/src/Camion/domain/entities"
	"github.com/vicpoo/API_recolecta/src/Camion/domain/ports"
)

type GetCamionByIDUseCase struct {
	repo ports.ICamion
}

func NewGetCamionByIDUseCase(repo ports.ICamion) *GetCamionByIDUseCase {
	return &GetCamionByIDUseCase{
		repo: repo,
	}
}

func (uc *GetCamionByIDUseCase) Run(id int32) (*entities.Camion, error) {
	if id <= 0 {
		return nil, errors.New("id inválido")
	}

	return uc.repo.GetByID(id)
}
