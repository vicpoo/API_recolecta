// GetReporteConductorByIdUseCase.go
package application

import (
	repositories "github.com/vicpoo/API_recolecta/src/reporte_conductor/domain"
	"github.com/vicpoo/API_recolecta/src/reporte_conductor/domain/entities"
)

type GetReporteConductorByIdUseCase struct {
	repo repositories.IReporteConductor
}

func NewGetReporteConductorByIdUseCase(repo repositories.IReporteConductor) *GetReporteConductorByIdUseCase {
	return &GetReporteConductorByIdUseCase{repo: repo}
}

func (uc *GetReporteConductorByIdUseCase) Run(id int32) (*entities.ReporteConductor, error) {
	return uc.repo.GetByID(id)
}