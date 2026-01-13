// seguimiento_falla_critica_repository.go
package domain

import (
	"github.com/vicpoo/API_recolecta/src/seguimiento_falla_critica/domain/entities"
)

type ISeguimientoFallaCritica interface {
	// Operaciones CRUD básicas
	Save(seguimiento *entities.SeguimientoFallaCritica) error
	Update(seguimiento *entities.SeguimientoFallaCritica) error
	Delete(id int32) error
	GetAll() ([]entities.SeguimientoFallaCritica, error)
	GetByID(id int32) (*entities.SeguimientoFallaCritica, error)
	
	// Métodos específicos para SeguimientoFallaCritica
	GetByFallaID(fallaID int32) ([]entities.SeguimientoFallaCritica, error)
	GetByFechaRange(fechaInicio, fechaFin string) ([]entities.SeguimientoFallaCritica, error)
}