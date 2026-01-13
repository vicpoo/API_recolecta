// reporte_mantenimiento_generado_repository.go
package domain

import (
	"github.com/vicpoo/API_recolecta/src/reporte_mantenimiento_generado/domain/entities"
)

type IReporteMantenimientoGenerado interface {
	// Operaciones CRUD básicas
	Save(reporte *entities.ReporteMantenimientoGenerado) error
	Update(reporte *entities.ReporteMantenimientoGenerado) error
	Delete(id int32) error
	GetAll() ([]entities.ReporteMantenimientoGenerado, error)
	GetByID(id int32) (*entities.ReporteMantenimientoGenerado, error)
	
	// Métodos específicos para ReporteMantenimientoGenerado
	GetByCoordinadorID(coordinadorID int32) ([]entities.ReporteMantenimientoGenerado, error)
	GetByFechaRange(fechaInicio, fechaFin string) ([]entities.ReporteMantenimientoGenerado, error)
	GetByFechaGeneracionRange(fechaInicio, fechaFin string) ([]entities.ReporteMantenimientoGenerado, error)
}