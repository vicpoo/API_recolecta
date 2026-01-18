// GetReportesMantenimientoGeneradoByFechaRangeUseCase.go
package application

import (
	repositories "github.com/vicpoo/API_recolecta/src/reporte_mantenimiento_generado/domain"
	"github.com/vicpoo/API_recolecta/src/reporte_mantenimiento_generado/domain/entities"
)

type GetReportesMantenimientoGeneradoByFechaRangeUseCase struct {
	repo repositories.IReporteMantenimientoGenerado
}

func NewGetReportesMantenimientoGeneradoByFechaRangeUseCase(repo repositories.IReporteMantenimientoGenerado) *GetReportesMantenimientoGeneradoByFechaRangeUseCase {
	return &GetReportesMantenimientoGeneradoByFechaRangeUseCase{repo: repo}
}

func (uc *GetReportesMantenimientoGeneradoByFechaRangeUseCase) Run(fechaInicio, fechaFin string) ([]entities.ReporteMantenimientoGenerado, error) {
	return uc.repo.GetByFechaRange(fechaInicio, fechaFin)
}