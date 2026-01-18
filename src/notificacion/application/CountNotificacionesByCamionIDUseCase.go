//CountNotificacionesByCamionIDUseCase.go
package application

import repositories "github.com/vicpoo/API_recolecta/src/notificacion/domain"

type CountNotificacionesByCamionIDUseCase struct {
	repo repositories.INotificacion
}

func NewCountNotificacionesByCamionIDUseCase(repo repositories.INotificacion) *CountNotificacionesByCamionIDUseCase {
	return &CountNotificacionesByCamionIDUseCase{repo: repo}
}

func (uc *CountNotificacionesByCamionIDUseCase) Run(camionID int32) (int64, error) {
	return uc.repo.CountByCamionID(camionID)
}