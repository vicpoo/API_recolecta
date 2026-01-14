package application

import (
	"github.com/vicpoo/API_recolecta/src/RegistroVaciado/domain/entities"
	"github.com/vicpoo/API_recolecta/src/RegistroVaciado/domain/ports"
)

type CreateRegistroVaciadoUseCase struct {
	repo ports.RegistroVaciadoRepository
}

func NewCreateRegistroVaciadoUseCase(repo ports.RegistroVaciadoRepository) *CreateRegistroVaciadoUseCase {
	return &CreateRegistroVaciadoUseCase{repo: repo}
}

func (uc *CreateRegistroVaciadoUseCase) Execute(registro *entities.RegistroVaciado) (*entities.RegistroVaciado, error) {
	return uc.repo.Save(registro)
}
