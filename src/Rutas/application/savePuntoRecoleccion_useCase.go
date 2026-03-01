package application

import (
	"github.com/vicpoo/API_recolecta/src/Rutas/domain/entities"
	"github.com/vicpoo/API_recolecta/src/Rutas/domain/ports"
)

type SavePuntoRecoleccionUseCase struct {
	repo ports.IPuntoRecoleccion
}

func NewSavePuntoRecoleccionUseCase(repo ports.IPuntoRecoleccion) *SavePuntoRecoleccionUseCase {
	return &SavePuntoRecoleccionUseCase{repo: repo}
}

func (uc *SavePuntoRecoleccionUseCase) Execute(p *entities.PuntoRecoleccion) (*entities.PuntoRecoleccion, error) {
	return uc.repo.Save(p)
}
