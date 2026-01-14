package application

import (
	"github.com/vicpoo/API_recolecta/src/Ruta/domain/entities"
	"github.com/vicpoo/API_recolecta/src/Ruta/domain/ports"
)


type UpdateRutaUseCase struct {
	repo ports.IRuta
}

func NewUpdateRutaUseCase(repo ports.IRuta) *UpdateRutaUseCase {
	return &UpdateRutaUseCase{repo}
}

func (uc *UpdateRutaUseCase) Run(ruta *entities.Ruta) error {
	return uc.repo.Update(ruta)
}
