//CountNotificacionesActivasByUsuarioIDUseCase.go
package application

import repositories "github.com/vicpoo/API_recolecta/src/notificacion/domain"

type CountNotificacionesActivasByUsuarioIDUseCase struct {
	repo repositories.INotificacion
}

func NewCountNotificacionesActivasByUsuarioIDUseCase(repo repositories.INotificacion) *CountNotificacionesActivasByUsuarioIDUseCase {
	return &CountNotificacionesActivasByUsuarioIDUseCase{repo: repo}
}

func (uc *CountNotificacionesActivasByUsuarioIDUseCase) Run(usuarioID int32) (int64, error) {
	return uc.repo.CountActivasByUsuarioID(usuarioID)
}