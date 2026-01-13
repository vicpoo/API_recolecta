//GetIncidenciaByIDUseCase.go
package application

import (
	repositories "github.com/vicpoo/API_recolecta/src/incidencia/domain"
	"github.com/vicpoo/API_recolecta/src/incidencia/domain/entities"
)

type GetIncidenciaByIDUseCase struct {
	repo repositories.IIncidencia
}

func NewGetIncidenciaByIDUseCase(repo repositories.IIncidencia) *GetIncidenciaByIDUseCase {
	return &GetIncidenciaByIDUseCase{repo: repo}
}

func (uc *GetIncidenciaByIDUseCase) Run(id int32) (*entities.Incidencia, error) {
	return uc.repo.GetByID(id)
}