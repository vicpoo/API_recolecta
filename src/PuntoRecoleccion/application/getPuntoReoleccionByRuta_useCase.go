package application

import (
	"github.com/vicpoo/API_recolecta/src/PuntoRecoleccion/domain/entities"
	"github.com/vicpoo/API_recolecta/src/PuntoRecoleccion/domain/ports"
)

type GetPuntoRecoleccionByRutaUseCase struct {
	repo ports.IPuntoRecoleccion
}

func NewGetPuntoRecoleccionByRutaUseCase(repo ports.IPuntoRecoleccion) *GetPuntoRecoleccionByRutaUseCase {
	return &GetPuntoRecoleccionByRutaUseCase{
		repo: repo,
	}
}

func (uc *GetPuntoRecoleccionByRutaUseCase) Execute(rutaId int32) ([]entities.PuntoRecoleccion, error) {
	return uc.repo.GetByRuta(rutaId)
}
