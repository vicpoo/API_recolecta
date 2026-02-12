package application

import (
	"github.com/vicpoo/API_recolecta/src/Rutas/domain/entities"
	"github.com/vicpoo/API_recolecta/src/Rutas/domain/ports"
)


type ListAllRutaUseCase struct {
	repo ports.IRuta
}

func NewListAllRutaUseCase(repo ports.IRuta) *ListAllRutaUseCase {
	return &ListAllRutaUseCase{repo}
}

func (uc *ListAllRutaUseCase) Run() ([]entities.Ruta, error) {
	return uc.repo.ListAll()
}
