// GetAllSeguimientosFallaCriticaUseCase.go
package application

import (
	repositories "github.com/vicpoo/API_recolecta/src/seguimiento_falla_critica/domain"
	"github.com/vicpoo/API_recolecta/src/seguimiento_falla_critica/domain/entities"
)

type GetAllSeguimientosFallaCriticaUseCase struct {
	repo repositories.ISeguimientoFallaCritica
}

func NewGetAllSeguimientosFallaCriticaUseCase(repo repositories.ISeguimientoFallaCritica) *GetAllSeguimientosFallaCriticaUseCase {
	return &GetAllSeguimientosFallaCriticaUseCase{repo: repo}
}

func (uc *GetAllSeguimientosFallaCriticaUseCase) Run() ([]entities.SeguimientoFallaCritica, error) {
	return uc.repo.GetAll()
}