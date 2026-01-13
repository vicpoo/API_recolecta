// GetReportesConductorByCamionIDUseCase.go
package application

import (
	repositories "github.com/vicpoo/API_recolecta/src/reporte_conductor/domain"
	"github.com/vicpoo/API_recolecta/src/reporte_conductor/domain/entities"
)

type GetReportesConductorByCamionIDUseCase struct {
	repo repositories.IReporteConductor
}

func NewGetReportesConductorByCamionIDUseCase(repo repositories.IReporteConductor) *GetReportesConductorByCamionIDUseCase {
	return &GetReportesConductorByCamionIDUseCase{repo: repo}
}

func (uc *GetReportesConductorByCamionIDUseCase) Run(camionID int32) ([]entities.ReporteConductor, error) {
	return uc.repo.GetByCamionID(camionID)
}