package application

import (
	"github.com/vicpoo/API_recolecta/src/RegistroVaciado/domain/entities"
	"github.com/vicpoo/API_recolecta/src/RegistroVaciado/domain/ports"
)

type ListAllRegistroVaciadoUseCase struct {
	repo ports.RegistroVaciadoRepository
}

func NewListAllRegistroVaciadoUseCase(repo ports.RegistroVaciadoRepository) *ListAllRegistroVaciadoUseCase {
	return &ListAllRegistroVaciadoUseCase{repo: repo}
}

func (uc *ListAllRegistroVaciadoUseCase) Execute() ([]entities.RegistroVaciado, error) {
	return uc.repo.ListAll()
}
