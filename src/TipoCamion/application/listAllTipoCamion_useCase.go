package application

import (
	"github.com/vicpoo/API_recolecta/src/TipoCamion/domain/entities"
	"github.com/vicpoo/API_recolecta/src/TipoCamion/domain/ports"
)

type ListAllTipoCamionUseCase struct {
	ITipoCamion ports.ITipoCamion
}

func NewListAllTipoCamion(ITipoCamion ports.ITipoCamion) *ListAllTipoCamionUseCase {
	return &ListAllTipoCamionUseCase{
		ITipoCamion: ITipoCamion,
	}
}

func (uc *ListAllTipoCamionUseCase) Run() ([]entities.TipoCamion, error) {
	tiposCamion, err := uc.ITipoCamion.ListAll()

	if err != nil {
		return nil, err
	}

	return tiposCamion, nil
}