// GetAnomaliasByChoferIDUseCase.go
package application

import (
	repositories "github.com/vicpoo/API_recolecta/src/anomalia/domain"
	"github.com/vicpoo/API_recolecta/src/anomalia/domain/entities"
)

type GetAnomaliasByChoferIDUseCase struct {
	repo repositories.IAnomalia
}

func NewGetAnomaliasByChoferIDUseCase(repo repositories.IAnomalia) *GetAnomaliasByChoferIDUseCase {
	return &GetAnomaliasByChoferIDUseCase{repo: repo}
}

func (uc *GetAnomaliasByChoferIDUseCase) Run(choferID int32) ([]entities.Anomalia, error) {
	return uc.repo.GetByChoferID(choferID)
}