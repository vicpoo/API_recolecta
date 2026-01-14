// CreateReporteFallaCriticaUseCase.go
package application

import (
	repositories "github.com/vicpoo/API_recolecta/src/reporte_falla_critica/domain"
	"github.com/vicpoo/API_recolecta/src/reporte_falla_critica/domain/entities"
)

type CreateReporteFallaCriticaUseCase struct {
	repo repositories.IReporteFallaCritica
}

func NewCreateReporteFallaCriticaUseCase(repo repositories.IReporteFallaCritica) *CreateReporteFallaCriticaUseCase {
	return &CreateReporteFallaCriticaUseCase{repo: repo}
}

func (uc *CreateReporteFallaCriticaUseCase) Run(reporte *entities.ReporteFallaCritica) (*entities.ReporteFallaCritica, error) {
	err := uc.repo.Save(reporte)
	if err != nil {
		return nil, err
	}
	return reporte, nil
}