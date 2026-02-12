package application

import (
	"errors"

	"github.com/vicpoo/API_recolecta/src/Rutas/domain/entities"
	"github.com/vicpoo/API_recolecta/src/Rutas/domain/ports"
)

type UpdateCamionUseCase struct {
	repo ports.ICamion
}

func NewUpdateCamionUseCase(repo ports.ICamion) *UpdateCamionUseCase {
	return &UpdateCamionUseCase{
		repo: repo,
	}
}

func (uc *UpdateCamionUseCase) Run(id int32, camion *entities.Camion) (*entities.Camion, error) {
	if id <= 0 {
		return nil, errors.New("id inválido")
	}

	return uc.repo.Update(id,camion)
}
