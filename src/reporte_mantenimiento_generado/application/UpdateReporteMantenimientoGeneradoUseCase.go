// UpdateReporteMantenimientoGeneradoUseCase.go
package application

import (
	repositories "github.com/vicpoo/API_recolecta/src/reporte_mantenimiento_generado/domain"
	"github.com/vicpoo/API_recolecta/src/reporte_mantenimiento_generado/domain/entities"
)

type UpdateReporteMantenimientoGeneradoUseCase struct {
	repo repositories.IReporteMantenimientoGenerado
}

func NewUpdateReporteMantenimientoGeneradoUseCase(repo repositories.IReporteMantenimientoGenerado) *UpdateReporteMantenimientoGeneradoUseCase {
	return &UpdateReporteMantenimientoGeneradoUseCase{repo: repo}
}

func (uc *UpdateReporteMantenimientoGeneradoUseCase) Run(reporte *entities.ReporteMantenimientoGenerado) (*entities.ReporteMantenimientoGenerado, error) {
	err := uc.repo.Update(reporte)
	if err != nil {
		return nil, err
	}
	return reporte, nil
}