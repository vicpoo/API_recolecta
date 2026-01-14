package application

import (
	"github.com/vicpoo/API_recolecta/src/PuntoRecoleccion/domain/entities"
	"github.com/vicpoo/API_recolecta/src/PuntoRecoleccion/domain/ports"
)

type UpdatePuntoRecoleccionUseCase struct {
	repo ports.IPuntoRecoleccion
}

func NewUpdatePuntoRecoleccionUseCase(repo ports.IPuntoRecoleccion) *UpdatePuntoRecoleccionUseCase {
	return &UpdatePuntoRecoleccionUseCase{repo: repo}
}

func (uc *UpdatePuntoRecoleccionUseCase) Execute(id int32, p *entities.PuntoRecoleccion) (*entities.PuntoRecoleccion, error) {
	return uc.repo.Update(id, p)
}
