package application

import "github.com/vicpoo/API_recolecta/src/Rutas/domain/ports"

type ListAllPuntoRecoleccionUseCase struct {
	repo ports.IPuntoRecoleccion
}

func NewListAllPuntoRecoleccionUseCase(repo ports.IPuntoRecoleccion) *ListAllPuntoRecoleccionUseCase {
	return &ListAllPuntoRecoleccionUseCase{repo: repo}
}

func (uc *ListAllPuntoRecoleccionUseCase) Execute() (interface{}, error) {
	return uc.repo.ListAll()
}
