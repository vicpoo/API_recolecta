//GetIncidenciasByConductorIDUseCase.go
package application

import (
	repositories "github.com/vicpoo/API_recolecta/src/incidencia/domain"
	"github.com/vicpoo/API_recolecta/src/incidencia/domain/entities"
)

type GetIncidenciasByConductorIDUseCase struct {
	repo repositories.IIncidencia
}

func NewGetIncidenciasByConductorIDUseCase(repo repositories.IIncidencia) *GetIncidenciasByConductorIDUseCase {
	return &GetIncidenciasByConductorIDUseCase{repo: repo}
}

func (uc *GetIncidenciasByConductorIDUseCase) Run(conductorID int32) ([]entities.Incidencia, error) {
	return uc.repo.GetByConductorID(conductorID)
}