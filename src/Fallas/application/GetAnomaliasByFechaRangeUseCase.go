// GetAnomaliasByFechaRangeUseCase.go
package application

import (
	repositories "github.com/vicpoo/API_recolecta/src/Fallas/domain"
	"github.com/vicpoo/API_recolecta/src/Fallas/domain/entities"
)

type GetAnomaliasByFechaRangeUseCase struct {
	repo repositories.IAnomalia
}

func NewGetAnomaliasByFechaRangeUseCase(repo repositories.IAnomalia) *GetAnomaliasByFechaRangeUseCase {
	return &GetAnomaliasByFechaRangeUseCase{repo: repo}
}

func (uc *GetAnomaliasByFechaRangeUseCase) Run(fechaInicio, fechaFin string) ([]entities.Anomalia, error) {
	return uc.repo.GetByFechaRange(fechaInicio, fechaFin)
}