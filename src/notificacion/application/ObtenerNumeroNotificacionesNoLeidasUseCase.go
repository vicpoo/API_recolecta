//ObtenerNumeroNotificacionesNoLeidasUseCase.go
package application

import repositories "github.com/vicpoo/API_recolecta/src/notificacion/domain"

type ObtenerNumeroNotificacionesNoLeidasUseCase struct {
	repo repositories.INotificacion
}

func NewObtenerNumeroNotificacionesNoLeidasUseCase(repo repositories.INotificacion) *ObtenerNumeroNotificacionesNoLeidasUseCase {
	return &ObtenerNumeroNotificacionesNoLeidasUseCase{repo: repo}
}

func (uc *ObtenerNumeroNotificacionesNoLeidasUseCase) Run(usuarioID int32) (int64, error) {
	return uc.repo.CountActivasByUsuarioID(usuarioID)
}