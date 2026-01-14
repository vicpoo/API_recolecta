// GetAnomaliasByTipoAnomaliaUseCase.go
package application

import (
	repositories "github.com/vicpoo/API_recolecta/src/anomalia/domain"
	"github.com/vicpoo/API_recolecta/src/anomalia/domain/entities"
)

type GetAnomaliasByTipoAnomaliaUseCase struct {
	repo repositories.IAnomalia
}

func NewGetAnomaliasByTipoAnomaliaUseCase(repo repositories.IAnomalia) *GetAnomaliasByTipoAnomaliaUseCase {
	return &GetAnomaliasByTipoAnomaliaUseCase{repo: repo}
}

func (uc *GetAnomaliasByTipoAnomaliaUseCase) Run(tipoAnomalia string) ([]entities.Anomalia, error) {
	return uc.repo.GetByTipoAnomalia(tipoAnomalia)
}