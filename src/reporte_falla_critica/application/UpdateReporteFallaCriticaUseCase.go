// UpdateReporteFallaCriticaUseCase.go
package application

import (
	repositories "github.com/vicpoo/API_recolecta/src/reporte_falla_critica/domain"
	"github.com/vicpoo/API_recolecta/src/reporte_falla_critica/domain/entities"
)

type UpdateReporteFallaCriticaUseCase struct {
	repo repositories.IReporteFallaCritica
}

func NewUpdateReporteFallaCriticaUseCase(repo repositories.IReporteFallaCritica) *UpdateReporteFallaCriticaUseCase {
	return &UpdateReporteFallaCriticaUseCase{repo: repo}
}

func (uc *UpdateReporteFallaCriticaUseCase) Run(reporte *entities.ReporteFallaCritica) (*entities.ReporteFallaCritica, error) {
	// Primero obtenemos el reporte actual para preservar el campo eliminado
	reporteActual, err := uc.repo.GetByID(reporte.GetFallaID())
	if err != nil {
		return nil, err
	}
	
	// Preservamos el campo eliminado del reporte actual
	// No permitimos que se actualice el campo eliminado a través de este método
	reporte.SetEliminado(reporteActual.GetEliminado())
	
	err = uc.repo.Update(reporte)
	if err != nil {
		return nil, err
	}
	return reporte, nil
}