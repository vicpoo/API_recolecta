package application

import (
	"errors"

	"github.com/vicpoo/API_recolecta/src/Camion/domain/ports"
)

type DeleteTipoCamionUseCase struct {
	ITipoCamion ports.ITipoCamion
}

func NewDeleteTipoCamionUseCase(
	ITipoCamion ports.ITipoCamion,
) *DeleteTipoCamionUseCase {
	return &DeleteTipoCamionUseCase{
		ITipoCamion: ITipoCamion,
	}
}

func (uc *DeleteTipoCamionUseCase) Run(id int32) error {
	if id <= 0 {
		return errors.New("id inválido")
	}

	err := uc.ITipoCamion.Delete(id)
	if err != nil {
		return err
	}

	return nil
}
