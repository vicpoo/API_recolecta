package application

import (
	"github.com/vicpoo/API_recolecta/src/EstadoCamion/domain/entities"
	"github.com/vicpoo/API_recolecta/src/EstadoCamion/domain/ports"
)


type UpdateEstadoCamionUseCase struct {
	IEstadoCamion ports.IEstadoCamion
}

func NewUpdateEstadoCamionUseCase(IEstadoCamion ports.IEstadoCamion) UpdateEstadoCamionUseCase {
	return UpdateEstadoCamionUseCase{
		IEstadoCamion: IEstadoCamion,
	}
}

func (uc *UpdateEstadoCamionUseCase) Run(id int32, estadoCamion *entities.EstadoCamion) (*entities.EstadoCamion, error) {
	estadoCamionUpdated, err := uc.IEstadoCamion.Update(id, estadoCamion)

	if err != nil {
		return nil, err
	}

	return estadoCamionUpdated, nil
}