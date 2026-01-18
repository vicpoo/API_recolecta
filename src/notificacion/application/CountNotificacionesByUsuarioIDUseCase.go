//CountNotificacionesByUsuarioIDUseCase.go
package application

import repositories "github.com/vicpoo/API_recolecta/src/notificacion/domain"

type CountNotificacionesByUsuarioIDUseCase struct {
	repo repositories.INotificacion
}

func NewCountNotificacionesByUsuarioIDUseCase(repo repositories.INotificacion) *CountNotificacionesByUsuarioIDUseCase {
	return &CountNotificacionesByUsuarioIDUseCase{repo: repo}
}

func (uc *CountNotificacionesByUsuarioIDUseCase) Run(usuarioID int32) (int64, error) {
	return uc.repo.CountByUsuarioID(usuarioID)
}