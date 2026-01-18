// GetReportesConductorByFechaRangeUseCase.go
package application

import (
	repositories "github.com/vicpoo/API_recolecta/src/reporte_conductor/domain"
	"github.com/vicpoo/API_recolecta/src/reporte_conductor/domain/entities"
)

type GetReportesConductorByFechaRangeUseCase struct {
	repo repositories.IReporteConductor
}

func NewGetReportesConductorByFechaRangeUseCase(repo repositories.IReporteConductor) *GetReportesConductorByFechaRangeUseCase {
	return &GetReportesConductorByFechaRangeUseCase{repo: repo}
}

func (uc *GetReportesConductorByFechaRangeUseCase) Run(fechaInicio, fechaFin string) ([]entities.ReporteConductor, error) {
	return uc.repo.GetByFechaRange(fechaInicio, fechaFin)
}