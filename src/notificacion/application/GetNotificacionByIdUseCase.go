//GetNotificacionByIdUseCase.go
package application

import (
	repositories "github.com/vicpoo/API_recolecta/src/notificacion/domain"
	"github.com/vicpoo/API_recolecta/src/notificacion/domain/entities"
)

type GetNotificacionByIdUseCase struct {
	repo repositories.INotificacion
}

func NewGetNotificacionByIdUseCase(repo repositories.INotificacion) *GetNotificacionByIdUseCase {
	return &GetNotificacionByIdUseCase{repo: repo}
}

func (uc *GetNotificacionByIdUseCase) Run(id int32) (*entities.Notificacion, error) {
	return uc.repo.GetByID(id)
}