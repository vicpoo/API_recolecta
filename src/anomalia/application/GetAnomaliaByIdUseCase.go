// GetAnomaliaByIdUseCase.go
package application

import (
	repositories "github.com/vicpoo/API_recolecta/src/anomalia/domain"
	"github.com/vicpoo/API_recolecta/src/anomalia/domain/entities"
)

type GetAnomaliaByIdUseCase struct {
	repo repositories.IAnomalia
}

func NewGetAnomaliaByIdUseCase(repo repositories.IAnomalia) *GetAnomaliaByIdUseCase {
	return &GetAnomaliaByIdUseCase{repo: repo}
}

func (uc *GetAnomaliaByIdUseCase) Run(id int32) (*entities.Anomalia, error) {
	return uc.repo.GetByID(id)
}