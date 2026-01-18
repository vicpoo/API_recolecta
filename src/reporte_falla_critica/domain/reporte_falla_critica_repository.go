// reporte_falla_critica_repository.go
package domain

import (
	"github.com/vicpoo/API_recolecta/src/reporte_falla_critica/domain/entities"
)

type IReporteFallaCritica interface {
	// Operaciones CRUD básicas
	Save(reporteFallaCritica *entities.ReporteFallaCritica) error
	Update(reporteFallaCritica *entities.ReporteFallaCritica) error
	Delete(id int32) error
	GetAll() ([]entities.ReporteFallaCritica, error)
	GetByID(id int32) (*entities.ReporteFallaCritica, error)
	
	// Métodos específicos para ReporteFallaCritica
	GetByCamionID(camionID int32) ([]entities.ReporteFallaCritica, error)
	GetByConductorID(conductorID int32) ([]entities.ReporteFallaCritica, error)
	GetByFechaRange(fechaInicio, fechaFin string) ([]entities.ReporteFallaCritica, error)
}