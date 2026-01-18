// CreateSeguimientoFallaCriticaUseCase.go
package application

import (
	repositories "github.com/vicpoo/API_recolecta/src/seguimiento_falla_critica/domain"
	"github.com/vicpoo/API_recolecta/src/seguimiento_falla_critica/domain/entities"
)

type CreateSeguimientoFallaCriticaUseCase struct {
	repo repositories.ISeguimientoFallaCritica
}

func NewCreateSeguimientoFallaCriticaUseCase(repo repositories.ISeguimientoFallaCritica) *CreateSeguimientoFallaCriticaUseCase {
	return &CreateSeguimientoFallaCriticaUseCase{repo: repo}
}

func (uc *CreateSeguimientoFallaCriticaUseCase) Run(seguimiento *entities.SeguimientoFallaCritica) (*entities.SeguimientoFallaCritica, error) {
	err := uc.repo.Save(seguimiento)
	if err != nil {
		return nil, err
	}
	return seguimiento, nil
}