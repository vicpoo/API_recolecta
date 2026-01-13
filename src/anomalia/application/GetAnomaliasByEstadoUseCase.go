// GetAnomaliasByEstadoUseCase.go
package application

import (
	repositories "github.com/vicpoo/API_recolecta/src/anomalia/domain"
	"github.com/vicpoo/API_recolecta/src/anomalia/domain/entities"
)

type GetAnomaliasByEstadoUseCase struct {
	repo repositories.IAnomalia
}

func NewGetAnomaliasByEstadoUseCase(repo repositories.IAnomalia) *GetAnomaliasByEstadoUseCase {
	return &GetAnomaliasByEstadoUseCase{repo: repo}
}

func (uc *GetAnomaliasByEstadoUseCase) Run(estado string) ([]entities.Anomalia, error) {
	return uc.repo.GetByEstado(estado)
}