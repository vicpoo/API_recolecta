//GetNotificacionesByUsuarioYTipoUseCase.go
package application

import (
	repositories "github.com/vicpoo/API_recolecta/src/notificacion/domain"
	"github.com/vicpoo/API_recolecta/src/notificacion/domain/entities"
)

type GetNotificacionesByUsuarioYTipoUseCase struct {
	repo repositories.INotificacion
}

func NewGetNotificacionesByUsuarioYTipoUseCase(repo repositories.INotificacion) *GetNotificacionesByUsuarioYTipoUseCase {
	return &GetNotificacionesByUsuarioYTipoUseCase{repo: repo}
}

func (uc *GetNotificacionesByUsuarioYTipoUseCase) Run(usuarioID int32, tipo string) ([]entities.Notificacion, error) {
	return uc.repo.GetByUsuarioYTipo(usuarioID, tipo)
}