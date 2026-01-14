package application

import (
	"github.com/vicpoo/API_recolecta/src/EstadoCamion/domain/entities"
	"github.com/vicpoo/API_recolecta/src/EstadoCamion/domain/ports"
)

type ListAllEstadoCamionUseCase struct {
	IEstadoCamion ports.IEstadoCamion
}

func NewListAllEstadoCamionUseCase(IEstadoCamion ports.IEstadoCamion) *ListAllEstadoCamionUseCase {
	return &ListAllEstadoCamionUseCase{
		IEstadoCamion: IEstadoCamion,
	}
}

func (uc *ListAllEstadoCamionUseCase) Run() ([]entities.EstadoCamion, error) {
	estadosCamion, err := uc.IEstadoCamion.ListAll()

	if err != nil {
		return nil, err
	}

	return estadosCamion, nil
}