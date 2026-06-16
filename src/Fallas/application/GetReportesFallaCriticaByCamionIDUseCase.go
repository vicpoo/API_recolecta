// GetReportesFallaCriticaByCamionIDUseCase.go
package application

import (
	repositories "github.com/vicpoo/API_recolecta/src/Fallas/domain"
	"github.com/vicpoo/API_recolecta/src/Fallas/domain/entities"
)

type GetReportesFallaCriticaByCamionIDUseCase struct {
	repo repositories.IReporteFallaCritica
}

func NewGetReportesFallaCriticaByCamionIDUseCase(repo repositories.IReporteFallaCritica) *GetReportesFallaCriticaByCamionIDUseCase {
	return &GetReportesFallaCriticaByCamionIDUseCase{repo: repo}
}

func (uc *GetReportesFallaCriticaByCamionIDUseCase) Run(camionID int32) ([]entities.ReporteFallaCritica, error) {
	return uc.repo.GetByCamionID(camionID)
}