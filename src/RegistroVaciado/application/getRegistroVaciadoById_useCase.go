package application

import (
	"github.com/vicpoo/API_recolecta/src/RegistroVaciado/domain/entities"
	"github.com/vicpoo/API_recolecta/src/RegistroVaciado/domain/ports"
)

type GetRegistroVaciadoByIDUseCase struct {
	repo ports.RegistroVaciadoRepository
}

func NewGetRegistroVaciadoByIDUseCase(repo ports.RegistroVaciadoRepository) *GetRegistroVaciadoByIDUseCase {
	return &GetRegistroVaciadoByIDUseCase{repo: repo}
}

func (uc *GetRegistroVaciadoByIDUseCase) Execute(id int32) (*entities.RegistroVaciado, error) {
	return uc.repo.GetByID(id)
}
