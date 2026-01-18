// reporte_conductor_repository.go
package domain

import (
	"github.com/vicpoo/API_recolecta/src/reporte_conductor/domain/entities"
)

type IReporteConductor interface {
	// Operaciones CRUD básicas
	Save(reporteConductor *entities.ReporteConductor) error
	Update(reporteConductor *entities.ReporteConductor) error
	Delete(id int32) error
	GetAll() ([]entities.ReporteConductor, error)
	GetByID(id int32) (*entities.ReporteConductor, error)
	
	// Métodos específicos para ReporteConductor
	GetByCamionID(camionID int32) ([]entities.ReporteConductor, error)
	GetByConductorID(conductorID int32) ([]entities.ReporteConductor, error)
	GetByRutaID(rutaID int32) ([]entities.ReporteConductor, error)
	GetByFechaRange(fechaInicio, fechaFin string) ([]entities.ReporteConductor, error)
}