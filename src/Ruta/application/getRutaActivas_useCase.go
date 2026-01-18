package application

import (
	"github.com/vicpoo/API_recolecta/src/Ruta/domain/entities"
	"github.com/vicpoo/API_recolecta/src/Ruta/domain/ports"
)

type GetRutaActivasUseCase struct {
	repo ports.IRuta
}

func NewGetRutaActivasUseCase(repo ports.IRuta) *GetRutaActivasUseCase {
	return &GetRutaActivasUseCase{repo}
}

func (uc *GetRutaActivasUseCase) Run() ([]entities.Ruta, error) {
	return uc.repo.GetActivas()
}
