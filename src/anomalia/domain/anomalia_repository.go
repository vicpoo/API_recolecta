// anomalia_repository.go
package domain

import (
	"github.com/vicpoo/API_recolecta/src/anomalia/domain/entities"
)

type IAnomalia interface {
	// Operaciones CRUD básicas
	Save(anomalia *entities.Anomalia) error
	Update(anomalia *entities.Anomalia) error
	Delete(id int32) error
	GetAll() ([]entities.Anomalia, error)
	GetByID(id int32) (*entities.Anomalia, error)
	
	// Métodos específicos para Anomalia
	GetByPuntoID(puntoID int32) ([]entities.Anomalia, error)
	GetByChoferID(choferID int32) ([]entities.Anomalia, error)
	GetByEstado(estado string) ([]entities.Anomalia, error)
	GetByTipoAnomalia(tipoAnomalia string) ([]entities.Anomalia, error)
	GetByFechaRange(fechaInicio, fechaFin string) ([]entities.Anomalia, error)
}