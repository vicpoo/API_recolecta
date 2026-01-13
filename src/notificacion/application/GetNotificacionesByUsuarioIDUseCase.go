//GetNotificacionesByUsuarioIDUseCase.go
package application

import (
	repositories "github.com/vicpoo/API_recolecta/src/notificacion/domain"
	"github.com/vicpoo/API_recolecta/src/notificacion/domain/entities"
)

type GetNotificacionesByUsuarioIDUseCase struct {
	repo repositories.INotificacion
}

func NewGetNotificacionesByUsuarioIDUseCase(repo repositories.INotificacion) *GetNotificacionesByUsuarioIDUseCase {
	return &GetNotificacionesByUsuarioIDUseCase{repo: repo}
}

func (uc *GetNotificacionesByUsuarioIDUseCase) Run(usuarioID int32) ([]entities.Notificacion, error) {
	return uc.repo.GetByUsuarioID(usuarioID)
}