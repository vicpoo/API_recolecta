//MarcarAlertaAtendidaUseCase.go
package application

import repositories "github.com/vicpoo/API_recolecta/src/alerta_mantenimiento/domain"

type MarcarAlertaAtendidaUseCase struct {
	repo repositories.IAlertaMantenimiento
}

func NewMarcarAlertaAtendidaUseCase(repo repositories.IAlertaMantenimiento) *MarcarAlertaAtendidaUseCase {
	return &MarcarAlertaAtendidaUseCase{repo: repo}
}

func (uc *MarcarAlertaAtendidaUseCase) Run(id int32) error {
	return uc.repo.MarcarComoAtendida(id)
}
