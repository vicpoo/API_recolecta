package application

import "github.com/vicpoo/API_recolecta/src/PuntoRecoleccion/domain/ports"

type DeletePuntoRecoleccionUseCase struct {
	repo ports.IPuntoRecoleccion
}

func NewDeletePuntoRecoleccionUseCase(repo ports.IPuntoRecoleccion) *DeletePuntoRecoleccionUseCase {
	return &DeletePuntoRecoleccionUseCase{repo: repo}
}

func (uc *DeletePuntoRecoleccionUseCase) Execute(id int32) error {
	return uc.repo.Delete(id)
}
