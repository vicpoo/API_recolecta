// DeleteReporteFallaCriticaUseCase.go
package application

import (
	repositories "github.com/vicpoo/API_recolecta/src/reporte_falla_critica/domain"
)

type DeleteReporteFallaCriticaUseCase struct {
	repo repositories.IReporteFallaCritica
}

func NewDeleteReporteFallaCriticaUseCase(repo repositories.IReporteFallaCritica) *DeleteReporteFallaCriticaUseCase {
	return &DeleteReporteFallaCriticaUseCase{repo: repo}
}

func (uc *DeleteReporteFallaCriticaUseCase) Run(id int32) error {
	// Primero obtenemos el reporte para hacer soft delete
	reporte, err := uc.repo.GetByID(id)
	if err != nil {
		return err
	}
	
	// Marcamos como eliminado (soft delete)
	reporte.MarcarComoEliminado()
	
	// Actualizamos en la base de datos
	return uc.repo.Update(reporte)
}