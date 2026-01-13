// GetReportesFallaCriticaByConductorIDUseCase.go
package application

import (
	repositories "github.com/vicpoo/API_recolecta/src/reporte_falla_critica/domain"
	"github.com/vicpoo/API_recolecta/src/reporte_falla_critica/domain/entities"
)

type GetReportesFallaCriticaByConductorIDUseCase struct {
	repo repositories.IReporteFallaCritica
}

func NewGetReportesFallaCriticaByConductorIDUseCase(repo repositories.IReporteFallaCritica) *GetReportesFallaCriticaByConductorIDUseCase {
	return &GetReportesFallaCriticaByConductorIDUseCase{repo: repo}
}

func (uc *GetReportesFallaCriticaByConductorIDUseCase) Run(conductorID int32) ([]entities.ReporteFallaCritica, error) {
	return uc.repo.GetByConductorID(conductorID)
}