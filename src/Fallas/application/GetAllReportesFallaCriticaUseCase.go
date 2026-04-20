// GetAllReportesFallaCriticaUseCase.go
package application

import (
	repositories "github.com/vicpoo/API_recolecta/src/Fallas/domain"
	"github.com/vicpoo/API_recolecta/src/Fallas/domain/entities"
)

type GetAllReportesFallaCriticaUseCase struct {
	repo repositories.IReporteFallaCritica
}

func NewGetAllReportesFallaCriticaUseCase(repo repositories.IReporteFallaCritica) *GetAllReportesFallaCriticaUseCase {
	return &GetAllReportesFallaCriticaUseCase{repo: repo}
}

func (uc *GetAllReportesFallaCriticaUseCase) Run() ([]entities.ReporteFallaCritica, error) {
	return uc.repo.GetAll()
}