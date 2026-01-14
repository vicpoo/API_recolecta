// alerta_mantenimiento_repository.go
package domain

import (
	"time"
	"github.com/vicpoo/API_recolecta/src/alerta_mantenimiento/domain/entities"
)

type IAlertaMantenimiento interface {
	// Métodos CRUD básicos
	Save(alertaMantenimiento *entities.AlertaMantenimiento) error
	Update(alertaMantenimiento *entities.AlertaMantenimiento) error
	Delete(id int32) error
	GetAll() ([]entities.AlertaMantenimiento, error)
	GetByID(id int32) (*entities.AlertaMantenimiento, error)
	
	// Métodos específicos para alertas
	GetByCamionID(camionID int32) ([]entities.AlertaMantenimiento, error)
	GetByTipoMantenimientoID(tipoID int32) ([]entities.AlertaMantenimiento, error)
	GetPendientes() ([]entities.AlertaMantenimiento, error)
	GetAtendidas() ([]entities.AlertaMantenimiento, error)
	MarcarComoAtendida(id int32) error
	GetByFechaRange(fechaInicio, fechaFin time.Time) ([]entities.AlertaMantenimiento, error)
}