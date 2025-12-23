package application

import (
	"errors"

	"github.com/vicpoo/API_recolecta/src/TipoCamion/domain/entities"
	"github.com/vicpoo/API_recolecta/src/TipoCamion/domain/ports"
)

type GetTipoCamionByNameUseCase struct {
	ITipoCamion ports.ITipoCamion
}

func NewGetTipoCamionByNameUseCase(
	ITipoCamion ports.ITipoCamion,
) *GetTipoCamionByNameUseCase {
	return &GetTipoCamionByNameUseCase{
		ITipoCamion: ITipoCamion,
	}
}

func (uc *GetTipoCamionByNameUseCase) Run(nombre string) (*entities.TipoCamion, error) {
	tipoCamion, err := uc.ITipoCamion.GetByName(nombre)

	if err != nil {
		return nil, err
	}

	if tipoCamion == nil {
		return nil, errors.New("tipo camion no encontrado")
	}

	return tipoCamion, nil
}
