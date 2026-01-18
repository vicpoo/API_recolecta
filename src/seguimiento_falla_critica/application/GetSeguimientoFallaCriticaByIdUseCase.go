// GetSeguimientoFallaCriticaByIdUseCase.go
package application

import (
	repositories "github.com/vicpoo/API_recolecta/src/seguimiento_falla_critica/domain"
	"github.com/vicpoo/API_recolecta/src/seguimiento_falla_critica/domain/entities"
)

type GetSeguimientoFallaCriticaByIdUseCase struct {
	repo repositories.ISeguimientoFallaCritica
}

func NewGetSeguimientoFallaCriticaByIdUseCase(repo repositories.ISeguimientoFallaCritica) *GetSeguimientoFallaCriticaByIdUseCase {
	return &GetSeguimientoFallaCriticaByIdUseCase{repo: repo}
}

func (uc *GetSeguimientoFallaCriticaByIdUseCase) Run(id int32) (*entities.SeguimientoFallaCritica, error) {
	return uc.repo.GetByID(id)
}