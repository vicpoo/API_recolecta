//GetIncidenciaByIDUseCase.go
package application

import (
	repositories "github.com/vicpoo/API_recolecta/src/Fallas/domain"
	"github.com/vicpoo/API_recolecta/src/Fallas/domain/entities"
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