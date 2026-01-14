// DeleteSeguimientoFallaCriticaUseCase.go
package application

import repositories "github.com/vicpoo/API_recolecta/src/seguimiento_falla_critica/domain"

type DeleteSeguimientoFallaCriticaUseCase struct {
	repo repositories.ISeguimientoFallaCritica
}

func NewDeleteSeguimientoFallaCriticaUseCase(repo repositories.ISeguimientoFallaCritica) *DeleteSeguimientoFallaCriticaUseCase {
	return &DeleteSeguimientoFallaCriticaUseCase{repo: repo}
}

func (uc *DeleteSeguimientoFallaCriticaUseCase) Run(id int32) error {
	return uc.repo.Delete(id)
}