// UpdateSeguimientoFallaCriticaUseCase.go
package application

import (
	repositories "github.com/vicpoo/API_recolecta/src/seguimiento_falla_critica/domain"
	"github.com/vicpoo/API_recolecta/src/seguimiento_falla_critica/domain/entities"
)

type UpdateSeguimientoFallaCriticaUseCase struct {
	repo repositories.ISeguimientoFallaCritica
}

func NewUpdateSeguimientoFallaCriticaUseCase(repo repositories.ISeguimientoFallaCritica) *UpdateSeguimientoFallaCriticaUseCase {
	return &UpdateSeguimientoFallaCriticaUseCase{repo: repo}
}

func (uc *UpdateSeguimientoFallaCriticaUseCase) Run(seguimiento *entities.SeguimientoFallaCritica) (*entities.SeguimientoFallaCritica, error) {
	err := uc.repo.Update(seguimiento)
	if err != nil {
		return nil, err
	}
	return seguimiento, nil
}