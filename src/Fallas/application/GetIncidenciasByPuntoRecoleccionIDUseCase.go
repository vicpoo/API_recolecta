//GetIncidenciasByPuntoRecoleccionIDUseCase.go
package application

import (
	repositories "github.com/vicpoo/API_recolecta/src/Fallas/domain"
	"github.com/vicpoo/API_recolecta/src/Fallas/domain/entities"
)

type GetIncidenciasByPuntoRecoleccionIDUseCase struct {
	repo repositories.IIncidencia
}

func NewGetIncidenciasByPuntoRecoleccionIDUseCase(repo repositories.IIncidencia) *GetIncidenciasByPuntoRecoleccionIDUseCase {
	return &GetIncidenciasByPuntoRecoleccionIDUseCase{repo: repo}
}

func (uc *GetIncidenciasByPuntoRecoleccionIDUseCase) Run(puntoRecoleccionID int32) ([]entities.Incidencia, error) {
	return uc.repo.GetByPuntoRecoleccionID(puntoRecoleccionID)
}