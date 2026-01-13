// notificacion_repository.go
package domain

import (
	"github.com/vicpoo/API_recolecta/src/notificacion/domain/entities"
)

type INotificacion interface {
	// Operaciones CRUD básicas (5)
	Save(notificacion *entities.Notificacion) error
	Update(notificacion *entities.Notificacion) error
	Delete(id int32) error
	GetAll() ([]entities.Notificacion, error)
	GetByID(id int32) (*entities.Notificacion, error)
	
	// Métodos de consulta generales (10)
	GetByUsuarioID(usuarioID int32) ([]entities.Notificacion, error)
	GetActivasByUsuarioID(usuarioID int32) ([]entities.Notificacion, error)
	GetByTipo(tipo string) ([]entities.Notificacion, error)
	GetActivas() ([]entities.Notificacion, error)
	GetInactivas() ([]entities.Notificacion, error)
	GetByCreadoPor(creadoPor int32) ([]entities.Notificacion, error)
	GetByFechaRange(fechaInicio, fechaFin string) ([]entities.Notificacion, error)
	GetGlobales() ([]entities.Notificacion, error)
	
	// Métodos de consulta por relaciones específicas (3)
	GetByCamionID(camionID int32) ([]entities.Notificacion, error)
	GetByFallaID(fallaID int32) ([]entities.Notificacion, error)
	GetByMantenimientoID(mantenimientoID int32) ([]entities.Notificacion, error)
	
	// Métodos de consulta combinados (2)
	GetByCamionYTipo(camionID int32, tipo string) ([]entities.Notificacion, error)
	GetByUsuarioYTipo(usuarioID int32, tipo string) ([]entities.Notificacion, error)
	
	// Métodos de actualización (3)
	MarcarComoLeida(id int32) error
	MarcarComoActiva(id int32) error
	MarcarTodasComoLeidas(usuarioID int32) error
	
	// Métodos de conteo (4)
	CountActivasByUsuarioID(usuarioID int32) (int64, error)
	CountByUsuarioID(usuarioID int32) (int64, error)
	CountByTipo(tipo string) (int64, error)
	CountByCamionID(camionID int32) (int64, error)
	
	// Métodos para notificaciones relacionadas con acciones (3)
	CrearNotificacionFalla(usuarioID *int32, titulo string, mensaje string, camionID int32, fallaID int32, creadoPor *int32) error
	CrearNotificacionMantenimiento(usuarioID *int32, titulo string, mensaje string, camionID int32, mantenimientoID int32, creadoPor *int32) error
	CrearNotificacionEmergencia(usuarioID *int32, titulo string, mensaje string, camionID int32, creadoPor *int32) error
}