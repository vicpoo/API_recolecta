// UpdateAnomaliaUseCase.go
package application

import (
	repositories "github.com/vicpoo/API_recolecta/src/anomalia/domain"
	"github.com/vicpoo/API_recolecta/src/anomalia/domain/entities"
)

type UpdateAnomaliaUseCase struct {
	repo repositories.IAnomalia
}

func NewUpdateAnomaliaUseCase(repo repositories.IAnomalia) *UpdateAnomaliaUseCase {
	return &UpdateAnomaliaUseCase{repo: repo}
}

func (uc *UpdateAnomaliaUseCase) Run(anomalia *entities.Anomalia) (*entities.Anomalia, error) {
	err := uc.repo.Update(anomalia)
	if err != nil {
		return nil, err
	}
	return anomalia, nil
}