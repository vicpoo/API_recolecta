// GetReportesFallaCriticaByFechaRangeUseCase.go
package application

import (
	repositories "github.com/vicpoo/API_recolecta/src/reporte_falla_critica/domain"
	"github.com/vicpoo/API_recolecta/src/reporte_falla_critica/domain/entities"
)

type GetReportesFallaCriticaByFechaRangeUseCase struct {
	repo repositories.IReporteFallaCritica
}

func NewGetReportesFallaCriticaByFechaRangeUseCase(repo repositories.IReporteFallaCritica) *GetReportesFallaCriticaByFechaRangeUseCase {
	return &GetReportesFallaCriticaByFechaRangeUseCase{repo: repo}
}

func (uc *GetReportesFallaCriticaByFechaRangeUseCase) Run(fechaInicio, fechaFin string) ([]entities.ReporteFallaCritica, error) {
	return uc.repo.GetByFechaRange(fechaInicio, fechaFin)
}