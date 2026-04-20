// GetAllReportesConductorUseCase.go
package application

import (
	repositories "github.com/vicpoo/API_recolecta/src/Fallas/domain"
	"github.com/vicpoo/API_recolecta/src/Fallas/domain/entities"
)

type GetAllReportesConductorUseCase struct {
	repo repositories.IReporteConductor
}

func NewGetAllReportesConductorUseCase(repo repositories.IReporteConductor) *GetAllReportesConductorUseCase {
	return &GetAllReportesConductorUseCase{repo: repo}
}

func (uc *GetAllReportesConductorUseCase) Run() ([]entities.ReporteConductor, error) {
	return uc.repo.GetAll()
}