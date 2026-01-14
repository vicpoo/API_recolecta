package application

import (
	"github.com/vicpoo/API_recolecta/src/RegistroVaciado/domain/entities"
	"github.com/vicpoo/API_recolecta/src/RegistroVaciado/domain/ports"
)

type GetRegistroVaciadoByRellenoIDUseCase struct {
	repo ports.RegistroVaciadoRepository
}

func NewGetRegistroVaciadoByRellenoIDUseCase(repo ports.RegistroVaciadoRepository) *GetRegistroVaciadoByRellenoIDUseCase {
	return &GetRegistroVaciadoByRellenoIDUseCase{repo: repo}
}

func (uc *GetRegistroVaciadoByRellenoIDUseCase) Execute(rellenoID int32) ([]entities.RegistroVaciado, error) {
	return uc.repo.GetByRellenoID(rellenoID)
}
