package application

import (
	"github.com/vicpoo/API_recolecta/src/EstadoCamion/domain/entities"
	"github.com/vicpoo/API_recolecta/src/EstadoCamion/domain/ports"
)


type GetByIdEstadoCamionUseCase struct {
	IEstadoCamion ports.IEstadoCamion
}

func NewGetByIdEstadoCamionUseCase(IEstadoCamion ports.IEstadoCamion) *GetByIdEstadoCamionUseCase {
	return &GetByIdEstadoCamionUseCase{
		IEstadoCamion: IEstadoCamion,
	}
}

func (uc *GetByIdEstadoCamionUseCase) Run(id int32) (*entities.EstadoCamion, error){
	estadoCamion, err := uc.IEstadoCamion.GetById(id)

	if err != nil {
		return nil, err
	}

	return estadoCamion, nil
}