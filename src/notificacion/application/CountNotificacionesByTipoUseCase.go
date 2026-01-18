//CountNotificacionesByTipoUseCase.go
package application

import repositories "github.com/vicpoo/API_recolecta/src/notificacion/domain"

type CountNotificacionesByTipoUseCase struct {
	repo repositories.INotificacion
}

func NewCountNotificacionesByTipoUseCase(repo repositories.INotificacion) *CountNotificacionesByTipoUseCase {
	return &CountNotificacionesByTipoUseCase{repo: repo}
}

func (uc *CountNotificacionesByTipoUseCase) Run(tipo string) (int64, error) {
	return uc.repo.CountByTipo(tipo)
}