//GetNotificacionesActivasByUsuarioIDUseCase.go
package application

import (
	repositories "github.com/vicpoo/API_recolecta/src/notificacion/domain"
	"github.com/vicpoo/API_recolecta/src/notificacion/domain/entities"
)

type GetNotificacionesActivasByUsuarioIDUseCase struct {
	repo repositories.INotificacion
}

func NewGetNotificacionesActivasByUsuarioIDUseCase(repo repositories.INotificacion) *GetNotificacionesActivasByUsuarioIDUseCase {
	return &GetNotificacionesActivasByUsuarioIDUseCase{repo: repo}
}

func (uc *GetNotificacionesActivasByUsuarioIDUseCase) Run(usuarioID int32) ([]entities.Notificacion, error) {
	return uc.repo.GetActivasByUsuarioID(usuarioID)
}