// GetAnomaliasByPuntoIDUseCase.go
package application

import (
	repositories "github.com/vicpoo/API_recolecta/src/Fallas/domain"
	"github.com/vicpoo/API_recolecta/src/Fallas/domain/entities"
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