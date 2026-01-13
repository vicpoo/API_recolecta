// GetAllAnomaliasUseCase.go
package application

import (
	repositories "github.com/vicpoo/API_recolecta/src/anomalia/domain"
	"github.com/vicpoo/API_recolecta/src/anomalia/domain/entities"
)

type GetAllAnomaliasUseCase struct {
	repo repositories.IAnomalia
}

func NewGetAllAnomaliasUseCase(repo repositories.IAnomalia) *GetAllAnomaliasUseCase {
	return &GetAllAnomaliasUseCase{repo: repo}
}

func (uc *GetAllAnomaliasUseCase) Run() ([]entities.Anomalia, error) {
	return uc.repo.GetAll()
}