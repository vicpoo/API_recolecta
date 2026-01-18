//DeleteNotificacionUseCase.go
package application

import repositories "github.com/vicpoo/API_recolecta/src/notificacion/domain"

type DeleteNotificacionUseCase struct {
	repo repositories.INotificacion
}

func NewDeleteNotificacionUseCase(repo repositories.INotificacion) *DeleteNotificacionUseCase {
	return &DeleteNotificacionUseCase{repo: repo}
}

func (uc *DeleteNotificacionUseCase) Run(id int32) error {
	return uc.repo.Delete(id)
}