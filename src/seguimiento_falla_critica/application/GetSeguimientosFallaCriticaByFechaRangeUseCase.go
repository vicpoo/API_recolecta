// GetSeguimientosFallaCriticaByFechaRangeUseCase.go
package application

import (
	repositories "github.com/vicpoo/API_recolecta/src/seguimiento_falla_critica/domain"
	"github.com/vicpoo/API_recolecta/src/seguimiento_falla_critica/domain/entities"
)

type GetSeguimientosFallaCriticaByFechaRangeUseCase struct {
	repo repositories.ISeguimientoFallaCritica
}

func NewGetSeguimientosFallaCriticaByFechaRangeUseCase(repo repositories.ISeguimientoFallaCritica) *GetSeguimientosFallaCriticaByFechaRangeUseCase {
	return &GetSeguimientosFallaCriticaByFechaRangeUseCase{repo: repo}
}

func (uc *GetSeguimientosFallaCriticaByFechaRangeUseCase) Run(fechaInicio, fechaFin string) ([]entities.SeguimientoFallaCritica, error) {
	return uc.repo.GetByFechaRange(fechaInicio, fechaFin)
}