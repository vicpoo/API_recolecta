//CreateIncidenciaUseCase.go
package application

import (
	repositories "github.com/vicpoo/API_recolecta/src/incidencia/domain"
	"github.com/vicpoo/API_recolecta/src/incidencia/domain/entities"
)

type CreateIncidenciaUseCase struct {
	repo repositories.IIncidencia
}

func NewCreateIncidenciaUseCase(repo repositories.IIncidencia) *CreateIncidenciaUseCase {
	return &CreateIncidenciaUseCase{repo: repo}
}

func (uc *CreateIncidenciaUseCase) Run(incidencia *entities.Incidencia) (*entities.Incidencia, error) {
	err := uc.repo.Save(incidencia)
	if err != nil {
		return nil, err
	}
	return incidencia, nil
}