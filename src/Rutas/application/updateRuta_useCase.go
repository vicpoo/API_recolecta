package application

import (
	"github.com/vicpoo/API_recolecta/src/Rutas/domain/entities"
	"github.com/vicpoo/API_recolecta/src/Rutas/domain/ports"
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
