// GetAnomaliasByPuntoIDUseCase.go
package application

import (
	repositories "github.com/vicpoo/API_recolecta/src/anomalia/domain"
	"github.com/vicpoo/API_recolecta/src/anomalia/domain/entities"
)

type GetAnomaliasByPuntoIDUseCase struct {
	repo repositories.IAnomalia
}

func NewGetAnomaliasByPuntoIDUseCase(repo repositories.IAnomalia) *GetAnomaliasByPuntoIDUseCase {
	return &GetAnomaliasByPuntoIDUseCase{repo: repo}
}

func (uc *GetAnomaliasByPuntoIDUseCase) Run(puntoID int32) ([]entities.Anomalia, error) {
	return uc.repo.GetByPuntoID(puntoID)
}