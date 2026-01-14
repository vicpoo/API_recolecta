// GetReportesMantenimientoGeneradoByFechaGeneracionRangeUseCase.go
package application

import (
	repositories "github.com/vicpoo/API_recolecta/src/reporte_mantenimiento_generado/domain"
	"github.com/vicpoo/API_recolecta/src/reporte_mantenimiento_generado/domain/entities"
)

type GetReportesMantenimientoGeneradoByFechaGeneracionRangeUseCase struct {
	repo repositories.IReporteMantenimientoGenerado
}

func NewGetReportesMantenimientoGeneradoByFechaGeneracionRangeUseCase(repo repositories.IReporteMantenimientoGenerado) *GetReportesMantenimientoGeneradoByFechaGeneracionRangeUseCase {
	return &GetReportesMantenimientoGeneradoByFechaGeneracionRangeUseCase{repo: repo}
}

func (uc *GetReportesMantenimientoGeneradoByFechaGeneracionRangeUseCase) Run(fechaInicio, fechaFin string) ([]entities.ReporteMantenimientoGenerado, error) {
	return uc.repo.GetByFechaGeneracionRange(fechaInicio, fechaFin)
}