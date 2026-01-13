// DeleteReporteConductorUseCase.go
package application

import repositories "github.com/vicpoo/API_recolecta/src/reporte_conductor/domain"

type DeleteReporteConductorUseCase struct {
	repo repositories.IReporteConductor
}

func NewDeleteReporteConductorUseCase(repo repositories.IReporteConductor) *DeleteReporteConductorUseCase {
	return &DeleteReporteConductorUseCase{repo: repo}
}

func (uc *DeleteReporteConductorUseCase) Run(id int32) error {
	return uc.repo.Delete(id)
}