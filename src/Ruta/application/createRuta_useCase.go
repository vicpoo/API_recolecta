package application

import (
	"github.com/vicpoo/API_recolecta/src/Ruta/domain/entities"
	"github.com/vicpoo/API_recolecta/src/Ruta/domain/ports"
)


type CreateRutaUseCase struct {
	repo ports.IRuta
}

func NewCreateRutaUseCase(repo ports.IRuta) *CreateRutaUseCase {
	return &CreateRutaUseCase{repo}
}

func (uc *CreateRutaUseCase) Run(ruta *entities.Ruta) error {
	return uc.repo.Save(ruta)
}
