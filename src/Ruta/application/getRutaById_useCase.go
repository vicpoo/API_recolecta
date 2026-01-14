package application

import (
	"github.com/vicpoo/API_recolecta/src/Ruta/domain/entities"
	"github.com/vicpoo/API_recolecta/src/Ruta/domain/ports"
)


type GetRutaByIdUseCase struct {
	repo ports.IRuta
}

func NewGetRutaByIdUseCase(repo ports.IRuta) *GetRutaByIdUseCase {
	return &GetRutaByIdUseCase{repo}
}

func (uc *GetRutaByIdUseCase) Run(id int32) (*entities.Ruta, error) {
	return uc.repo.GetById(id)
}
