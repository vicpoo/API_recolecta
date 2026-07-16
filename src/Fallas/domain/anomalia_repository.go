// anomalia_repository.go
package domain

import (
	"github.com/vicpoo/API_recolecta/src/Fallas/domain/entities"
)

type IAnomalia interface {
	// Operaciones CRUD básicas
	Save(anomalia *entities.Anomalia) error
	Update(anomalia *entities.Anomalia) error
	Delete(id int32) error
	GetAll() ([]entities.Anomalia, error)
	GetByID(id int32) (*entities.Anomalia, error)

	// Métodos específicos para Anomalia (cubren los antiguos dominios
	// Incidencia, ReporteConductor, ReporteFallaCritica y
	// SeguimientoFallaCritica, ahora unificados por TipoAnomalia)
	GetByTipoAnomalia(tipoAnomalia string) ([]entities.Anomalia, error)
	GetByEstado(estado string) ([]entities.Anomalia, error)
	GetByPuntoID(puntoID int32) ([]entities.Anomalia, error)
	GetByConductorID(conductorID int32) ([]entities.Anomalia, error)
	GetByCamionID(camionID int32) ([]entities.Anomalia, error)
	GetByRutaID(rutaID int32) ([]entities.Anomalia, error)
	GetByReferenciaID(referenciaID int32) ([]entities.Anomalia, error)
	GetByFechaRange(fechaInicio, fechaFin string) ([]entities.Anomalia, error)
}
