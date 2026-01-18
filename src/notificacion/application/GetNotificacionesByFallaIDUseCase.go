//GetNotificacionesByFallaIDUseCase.go
package application

import (
	repositories "github.com/vicpoo/API_recolecta/src/notificacion/domain"
	"github.com/vicpoo/API_recolecta/src/notificacion/domain/entities"
)

type GetNotificacionesByFallaIDUseCase struct {
	repo repositories.INotificacion
}

func NewGetNotificacionesByFallaIDUseCase(repo repositories.INotificacion) *GetNotificacionesByFallaIDUseCase {
	return &GetNotificacionesByFallaIDUseCase{repo: repo}
}

func (uc *GetNotificacionesByFallaIDUseCase) Run(fallaID int32) ([]entities.Notificacion, error) {
	return uc.repo.GetByFallaID(fallaID)
}