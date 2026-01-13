//MarcarNotificacionComoLeidaUseCase.go
package application

import repositories "github.com/vicpoo/API_recolecta/src/notificacion/domain"

type MarcarNotificacionComoLeidaUseCase struct {
	repo repositories.INotificacion
}

func NewMarcarNotificacionComoLeidaUseCase(repo repositories.INotificacion) *MarcarNotificacionComoLeidaUseCase {
	return &MarcarNotificacionComoLeidaUseCase{repo: repo}
}

func (uc *MarcarNotificacionComoLeidaUseCase) Run(id int32) error {
	return uc.repo.MarcarComoLeida(id)
}