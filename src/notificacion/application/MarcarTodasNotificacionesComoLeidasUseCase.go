//MarcarTodasNotificacionesComoLeidasUseCase.go
package application

import repositories "github.com/vicpoo/API_recolecta/src/notificacion/domain"

type MarcarTodasNotificacionesComoLeidasUseCase struct {
	repo repositories.INotificacion
}

func NewMarcarTodasNotificacionesComoLeidasUseCase(repo repositories.INotificacion) *MarcarTodasNotificacionesComoLeidasUseCase {
	return &MarcarTodasNotificacionesComoLeidasUseCase{repo: repo}
}

func (uc *MarcarTodasNotificacionesComoLeidasUseCase) Run(usuarioID int32) error {
	return uc.repo.MarcarTodasComoLeidas(usuarioID)
}