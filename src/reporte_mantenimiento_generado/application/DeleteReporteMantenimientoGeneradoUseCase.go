// DeleteReporteMantenimientoGeneradoUseCase.go
package application

import repositories "github.com/vicpoo/API_recolecta/src/reporte_mantenimiento_generado/domain"

type DeleteReporteMantenimientoGeneradoUseCase struct {
	repo repositories.IReporteMantenimientoGenerado
}

func NewDeleteReporteMantenimientoGeneradoUseCase(repo repositories.IReporteMantenimientoGenerado) *DeleteReporteMantenimientoGeneradoUseCase {
	return &DeleteReporteMantenimientoGeneradoUseCase{repo: repo}
}

func (uc *DeleteReporteMantenimientoGeneradoUseCase) Run(id int32) error {
	return uc.repo.Delete(id)
}