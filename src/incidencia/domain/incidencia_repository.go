// incidencia_repository.go
package domain

import (
	"github.com/vicpoo/API_recolecta/src/incidencia/domain/entities"
)

type IIncidencia interface {
	// Operaciones CRUD básicas
	Save(incidencia *entities.Incidencia) error
	Update(incidencia *entities.Incidencia) error
	Delete(id int32) error
	GetAll() ([]entities.Incidencia, error)
	GetByID(id int32) (*entities.Incidencia, error)
	
	// Métodos específicos para Incidencia
	GetByConductorID(conductorID int32) ([]entities.Incidencia, error)
	GetByPuntoRecoleccionID(puntoRecoleccionID int32) ([]entities.Incidencia, error)
	GetByFechaRange(fechaInicio, fechaFin string) ([]entities.Incidencia, error)
}