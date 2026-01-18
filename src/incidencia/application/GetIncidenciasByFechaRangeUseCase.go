//GetIncidenciasByFechaRangeUseCase.go
package application

import (
	repositories "github.com/vicpoo/API_recolecta/src/incidencia/domain"
	"github.com/vicpoo/API_recolecta/src/incidencia/domain/entities"
)

type GetIncidenciasByFechaRangeUseCase struct {
	repo repositories.IIncidencia
}

func NewGetIncidenciasByFechaRangeUseCase(repo repositories.IIncidencia) *GetIncidenciasByFechaRangeUseCase {
	return &GetIncidenciasByFechaRangeUseCase{repo: repo}
}

func (uc *GetIncidenciasByFechaRangeUseCase) Run(fechaInicio, fechaFin string) ([]entities.Incidencia, error) {
	return uc.repo.GetByFechaRange(fechaInicio, fechaFin)
}