//MarcarNotificacionComoActivaUseCase.go
package application

import repositories "github.com/vicpoo/API_recolecta/src/notificacion/domain"

type MarcarNotificacionComoActivaUseCase struct {
	repo repositories.INotificacion
}

func NewMarcarNotificacionComoActivaUseCase(repo repositories.INotificacion) *MarcarNotificacionComoActivaUseCase {
	return &MarcarNotificacionComoActivaUseCase{repo: repo}
}

func (uc *MarcarNotificacionComoActivaUseCase) Run(id int32) error {
	return uc.repo.MarcarComoActiva(id)
}