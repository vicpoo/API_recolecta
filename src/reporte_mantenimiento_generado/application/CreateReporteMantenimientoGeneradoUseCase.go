// CreateReporteMantenimientoGeneradoUseCase.go
package application

import (
	repositories "github.com/vicpoo/API_recolecta/src/reporte_mantenimiento_generado/domain"
	"github.com/vicpoo/API_recolecta/src/reporte_mantenimiento_generado/domain/entities"
)

type CreateReporteMantenimientoGeneradoUseCase struct {
	repo repositories.IReporteMantenimientoGenerado
}

func NewCreateReporteMantenimientoGeneradoUseCase(repo repositories.IReporteMantenimientoGenerado) *CreateReporteMantenimientoGeneradoUseCase {
	return &CreateReporteMantenimientoGeneradoUseCase{repo: repo}
}

func (uc *CreateReporteMantenimientoGeneradoUseCase) Run(reporte *entities.ReporteMantenimientoGenerado) (*entities.ReporteMantenimientoGenerado, error) {
	err := uc.repo.Save(reporte)
	if err != nil {
		return nil, err
	}
	return reporte, nil
}