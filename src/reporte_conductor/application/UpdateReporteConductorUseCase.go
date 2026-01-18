// UpdateReporteConductorUseCase.go
package application

import (
	repositories "github.com/vicpoo/API_recolecta/src/reporte_conductor/domain"
	"github.com/vicpoo/API_recolecta/src/reporte_conductor/domain/entities"
)

type UpdateReporteConductorUseCase struct {
	repo repositories.IReporteConductor
}

func NewUpdateReporteConductorUseCase(repo repositories.IReporteConductor) *UpdateReporteConductorUseCase {
	return &UpdateReporteConductorUseCase{repo: repo}
}

func (uc *UpdateReporteConductorUseCase) Run(reporte *entities.ReporteConductor) (*entities.ReporteConductor, error) {
	err := uc.repo.Update(reporte)
	if err != nil {
		return nil, err
	}
	return reporte, nil
}