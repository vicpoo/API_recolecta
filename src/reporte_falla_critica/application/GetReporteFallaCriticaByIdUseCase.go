// GetReporteFallaCriticaByIdUseCase.go
package application

import (
	repositories "github.com/vicpoo/API_recolecta/src/reporte_falla_critica/domain"
	"github.com/vicpoo/API_recolecta/src/reporte_falla_critica/domain/entities"
)

type GetReporteFallaCriticaByIdUseCase struct {
	repo repositories.IReporteFallaCritica
}

func NewGetReporteFallaCriticaByIdUseCase(repo repositories.IReporteFallaCritica) *GetReporteFallaCriticaByIdUseCase {
	return &GetReporteFallaCriticaByIdUseCase{repo: repo}
}

func (uc *GetReporteFallaCriticaByIdUseCase) Run(id int32) (*entities.ReporteFallaCritica, error) {
	return uc.repo.GetByID(id)
}