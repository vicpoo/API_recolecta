// GetReportesConductorByRutaIDUseCase.go
package application

import (
	repositories "github.com/vicpoo/API_recolecta/src/reporte_conductor/domain"
	"github.com/vicpoo/API_recolecta/src/reporte_conductor/domain/entities"
)

type GetReportesConductorByRutaIDUseCase struct {
	repo repositories.IReporteConductor
}

func NewGetReportesConductorByRutaIDUseCase(repo repositories.IReporteConductor) *GetReportesConductorByRutaIDUseCase {
	return &GetReportesConductorByRutaIDUseCase{repo: repo}
}

func (uc *GetReportesConductorByRutaIDUseCase) Run(rutaID int32) ([]entities.ReporteConductor, error) {
	return uc.repo.GetByRutaID(rutaID)
}