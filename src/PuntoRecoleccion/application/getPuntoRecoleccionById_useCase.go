package application

import "github.com/vicpoo/API_recolecta/src/PuntoRecoleccion/domain/ports"

type GetPuntoRecoleccionByIdUseCase struct {
	repo ports.IPuntoRecoleccion
}

func NewGetPuntoRecoleccionByIdUseCase(repo ports.IPuntoRecoleccion) *GetPuntoRecoleccionByIdUseCase {
	return &GetPuntoRecoleccionByIdUseCase{repo: repo}
}

func (uc *GetPuntoRecoleccionByIdUseCase) Execute(id int32) (interface{}, error) {
	return uc.repo.GetById(id)
}
