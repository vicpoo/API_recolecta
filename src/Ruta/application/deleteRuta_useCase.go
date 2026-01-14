package application

import "github.com/vicpoo/API_recolecta/src/Ruta/domain/ports"



type DeleteRutaUseCase struct {
	repo ports.IRuta
}

func NewDeleteRutaUseCase(repo ports.IRuta) *DeleteRutaUseCase {
	return &DeleteRutaUseCase{repo}
}

func (uc *DeleteRutaUseCase) Run(id int32) error {
	return uc.repo.Delete(id)
}
