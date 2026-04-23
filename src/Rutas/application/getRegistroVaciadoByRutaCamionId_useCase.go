package application

import (
	"github.com/vicpoo/API_recolecta/src/Rutas/domain/entities"
	"github.com/vicpoo/API_recolecta/src/Rutas/domain/ports"
)

type GetRegistroVaciadoByRutaCamionIDUseCase struct {
	repo ports.RegistroVaciadoRepository
}

func NewGetRegistroVaciadoByRutaCamionIDUseCase(repo ports.RegistroVaciadoRepository) *GetRegistroVaciadoByRutaCamionIDUseCase {
	return &GetRegistroVaciadoByRutaCamionIDUseCase{repo: repo}
}

func (uc *GetRegistroVaciadoByRutaCamionIDUseCase) Execute(rutaCamionID int32) ([]entities.RegistroVaciado, error) {
	return uc.repo.GetByRutaCamionID(rutaCamionID)
}
