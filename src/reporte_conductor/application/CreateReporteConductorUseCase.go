// CreateReporteConductorUseCase.go
package application

import (
	repositories "github.com/vicpoo/API_recolecta/src/reporte_conductor/domain"
	"github.com/vicpoo/API_recolecta/src/reporte_conductor/domain/entities"
)

type CreateReporteConductorUseCase struct {
	repo repositories.IReporteConductor
}

func NewCreateReporteConductorUseCase(repo repositories.IReporteConductor) *CreateReporteConductorUseCase {
	return &CreateReporteConductorUseCase{repo: repo}
}

func (uc *CreateReporteConductorUseCase) Run(reporte *entities.ReporteConductor) (*entities.ReporteConductor, error) {
	err := uc.repo.Save(reporte)
	if err != nil {
		return nil, err
	}
	return reporte, nil
}