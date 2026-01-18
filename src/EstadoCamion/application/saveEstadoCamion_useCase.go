package application

import (
	"errors"

	"github.com/vicpoo/API_recolecta/src/EstadoCamion/domain/entities"
	"github.com/vicpoo/API_recolecta/src/EstadoCamion/domain/ports"
)

type SaveEstadoCamionUseCase struct {
	repo ports.IEstadoCamion
}

func NewSaveEstadoCamionUseCase(repo ports.IEstadoCamion) *SaveEstadoCamionUseCase {
	return &SaveEstadoCamionUseCase{repo: repo}
}

func (uc *SaveEstadoCamionUseCase) Run(estado *entities.EstadoCamion) (*entities.EstadoCamion, error) {
	if estado.Observaciones == "" {
		return nil, errors.New("el nombre es obligatorio")
	}

	if estado.Estado == "" {
		return nil, errors.New("el estado es necesario")
	}
	
	return uc.repo.Save(estado)
}
