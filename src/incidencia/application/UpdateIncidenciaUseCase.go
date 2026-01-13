//UpdateIncidenciaUseCase.go
package application

import (
	repositories "github.com/vicpoo/API_recolecta/src/incidencia/domain"
	"github.com/vicpoo/API_recolecta/src/incidencia/domain/entities"
	"time"
)

type UpdateIncidenciaUseCase struct {
	repo repositories.IIncidencia
}

func NewUpdateIncidenciaUseCase(repo repositories.IIncidencia) *UpdateIncidenciaUseCase {
	return &UpdateIncidenciaUseCase{repo: repo}
}

func (uc *UpdateIncidenciaUseCase) Run(incidencia *entities.Incidencia) (*entities.Incidencia, error) {
	// Actualizamos el timestamp
	incidencia.SetUpdatedAt(time.Now())
	
	err := uc.repo.Update(incidencia)
	if err != nil {
		return nil, err
	}
	return incidencia, nil
}