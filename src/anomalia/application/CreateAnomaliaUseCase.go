// CreateAnomaliaUseCase.go
package application

import (
	repositories "github.com/vicpoo/API_recolecta/src/anomalia/domain"
	"github.com/vicpoo/API_recolecta/src/anomalia/domain/entities"
)

type CreateAnomaliaUseCase struct {
	repo repositories.IAnomalia
}

func NewCreateAnomaliaUseCase(repo repositories.IAnomalia) *CreateAnomaliaUseCase {
	return &CreateAnomaliaUseCase{repo: repo}
}

func (uc *CreateAnomaliaUseCase) Run(anomalia *entities.Anomalia) (*entities.Anomalia, error) {
	err := uc.repo.Save(anomalia)
	if err != nil {
		return nil, err
	}
	return anomalia, nil
}