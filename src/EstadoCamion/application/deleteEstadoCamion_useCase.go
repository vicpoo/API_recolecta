package application

import (

	"github.com/vicpoo/API_recolecta/src/EstadoCamion/domain/ports"
)

type DeleteEstadoCamionUseCase struct {
	repo ports.IEstadoCamion
}

func NewDeleteEstadoCamionUseCase(repo ports.IEstadoCamion) *DeleteEstadoCamionUseCase {
	return &DeleteEstadoCamionUseCase{
		repo: repo,
	}
}

func (uc *DeleteEstadoCamionUseCase) Run(id int32) error {
 	return uc.repo.Delete(id)
}