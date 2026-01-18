// GetReporteMantenimientoGeneradoByIdUseCase.go
package application

import (
	repositories "github.com/vicpoo/API_recolecta/src/reporte_mantenimiento_generado/domain"
	"github.com/vicpoo/API_recolecta/src/reporte_mantenimiento_generado/domain/entities"
)

type GetReporteMantenimientoGeneradoByIdUseCase struct {
	repo repositories.IReporteMantenimientoGenerado
}

func NewGetReporteMantenimientoGeneradoByIdUseCase(repo repositories.IReporteMantenimientoGenerado) *GetReporteMantenimientoGeneradoByIdUseCase {
	return &GetReporteMantenimientoGeneradoByIdUseCase{repo: repo}
}

func (uc *GetReporteMantenimientoGeneradoByIdUseCase) Run(id int32) (*entities.ReporteMantenimientoGenerado, error) {
	return uc.repo.GetByID(id)
}