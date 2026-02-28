// GetAllReportesMantenimientoGeneradoUseCase.go
package application

import (
	repositories "github.com/vicpoo/API_recolecta/src/Mantenimiento/domain"
	"github.com/vicpoo/API_recolecta/src/Mantenimiento/domain/entities"
)

type GetAllReportesMantenimientoGeneradoUseCase struct {
	repo repositories.IReporteMantenimientoGenerado
}

func NewGetAllReportesMantenimientoGeneradoUseCase(repo repositories.IReporteMantenimientoGenerado) *GetAllReportesMantenimientoGeneradoUseCase {
	return &GetAllReportesMantenimientoGeneradoUseCase{repo: repo}
}

func (uc *GetAllReportesMantenimientoGeneradoUseCase) Run() ([]entities.ReporteMantenimientoGenerado, error) {
	return uc.repo.GetAll()
}