// GetReportesConductorByConductorIDUseCase.go
package application

import (
	repositories "github.com/vicpoo/API_recolecta/src/reporte_conductor/domain"
	"github.com/vicpoo/API_recolecta/src/reporte_conductor/domain/entities"
)

type GetReportesConductorByConductorIDUseCase struct {
	repo repositories.IReporteConductor
}

func NewGetReportesConductorByConductorIDUseCase(repo repositories.IReporteConductor) *GetReportesConductorByConductorIDUseCase {
	return &GetReportesConductorByConductorIDUseCase{repo: repo}
}

func (uc *GetReportesConductorByConductorIDUseCase) Run(conductorID int32) ([]entities.ReporteConductor, error) {
	return uc.repo.GetByConductorID(conductorID)
}