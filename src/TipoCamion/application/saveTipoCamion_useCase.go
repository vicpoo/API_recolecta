package application

import (
	"errors"
	"strings"

	"github.com/vicpoo/API_recolecta/src/TipoCamion/domain/entities"
	"github.com/vicpoo/API_recolecta/src/TipoCamion/domain/ports"
)

type SaveTipoCamionUseCase struct {
	ITipoCamion ports.ITipoCamion
}

func NewSaveTipoCamionUseCase(ITipoCamion ports.ITipoCamion) *SaveTipoCamionUseCase {
	return &SaveTipoCamionUseCase{
		ITipoCamion: ITipoCamion,
	}
}

func (uc *SaveTipoCamionUseCase) Run(tCamion *entities.TipoCamion) (*entities.TipoCamion, error) {

	// ===== VALIDACIONES =====

	tCamion.Nombre = strings.TrimSpace(tCamion.Nombre)
	tCamion.Descripcion = strings.TrimSpace(tCamion.Descripcion)

	if tCamion.Nombre == "" {
		return nil, errors.New("el nombre del tipo de camión es obligatorio")
	}

	if len(tCamion.Nombre) < 3 {
		return nil, errors.New("el nombre debe tener al menos 3 caracteres")
	}

	if len(tCamion.Nombre) > 100 {
		return nil, errors.New("el nombre no puede exceder 100 caracteres")
	}

	if len(tCamion.Descripcion) > 255 {
		return nil, errors.New("la descripción no puede exceder 255 caracteres")
	}

	// ===== PERSISTENCIA =====

	tipoCamion, err := uc.ITipoCamion.Save(tCamion)
	if err != nil {
		return nil, err
	}

	return tipoCamion, nil
}
