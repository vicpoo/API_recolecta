//GetAllNotificacionesUseCase.go
package application

import (
	repositories "github.com/vicpoo/API_recolecta/src/notificacion/domain"
	"github.com/vicpoo/API_recolecta/src/notificacion/domain/entities"
)

type GetAllNotificacionesUseCase struct {
	repo repositories.INotificacion
}

func NewGetAllNotificacionesUseCase(repo repositories.INotificacion) *GetAllNotificacionesUseCase {
	return &GetAllNotificacionesUseCase{repo: repo}
}

func (uc *GetAllNotificacionesUseCase) Run() ([]entities.Notificacion, error) {
	return uc.repo.GetAll()
}