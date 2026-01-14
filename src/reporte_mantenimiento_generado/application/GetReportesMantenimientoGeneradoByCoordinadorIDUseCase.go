// GetReportesMantenimientoGeneradoByCoordinadorIDUseCase.go
package application

import (
	repositories "github.com/vicpoo/API_recolecta/src/reporte_mantenimiento_generado/domain"
	"github.com/vicpoo/API_recolecta/src/reporte_mantenimiento_generado/domain/entities"
)

type GetReportesMantenimientoGeneradoByCoordinadorIDUseCase struct {
	repo repositories.IReporteMantenimientoGenerado
}

func NewGetReportesMantenimientoGeneradoByCoordinadorIDUseCase(repo repositories.IReporteMantenimientoGenerado) *GetReportesMantenimientoGeneradoByCoordinadorIDUseCase {
	return &GetReportesMantenimientoGeneradoByCoordinadorIDUseCase{repo: repo}
}

func (uc *GetReportesMantenimientoGeneradoByCoordinadorIDUseCase) Run(coordinadorID int32) ([]entities.ReporteMantenimientoGenerado, error) {
	return uc.repo.GetByCoordinadorID(coordinadorID)
}