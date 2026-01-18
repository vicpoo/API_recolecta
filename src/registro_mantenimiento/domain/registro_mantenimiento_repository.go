// registro_mantenimiento_repository.go
package domain

import (
	"github.com/vicpoo/API_recolecta/src/registro_mantenimiento/domain/entities"
)

type IRegistroMantenimiento interface {
	// Operaciones CRUD básicas
	Save(registroMantenimiento *entities.RegistroMantenimiento) error
	Update(registroMantenimiento *entities.RegistroMantenimiento) error
	Delete(id int32) error
	GetAll() ([]entities.RegistroMantenimiento, error)
	GetByID(id int32) (*entities.RegistroMantenimiento, error)
	
	// Métodos específicos para RegistroMantenimiento
	GetByCamionID(camionID int32) ([]entities.RegistroMantenimiento, error)
	GetByAlertaID(alertaID int32) (*entities.RegistroMantenimiento, error)
	GetByCoordinadorID(coordinadorID int32) ([]entities.RegistroMantenimiento, error)
	GetByFechaRange(fechaInicio, fechaFin string) ([]entities.RegistroMantenimiento, error)
}