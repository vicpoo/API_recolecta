package ports

import "github.com/vicpoo/API_recolecta/src/HistorialAsignacionCamion/domain/entities"

type IHistorialAsignacionCamion interface {
	Save(historial *entities.HistorialAsignacionCamion) (*entities.HistorialAsignacionCamion, error)
	GetById(id int32) (*entities.HistorialAsignacionCamion, error)
	ListAll() ([]entities.HistorialAsignacionCamion, error)
	Update(id int32, historial *entities.HistorialAsignacionCamion) (*entities.HistorialAsignacionCamion, error)
	Delete(id int32) error
	GetByCamionId(camionId int32) ([]entities.HistorialAsignacionCamion, error)
	GetByChoferId(choferId int32) ([]entities.HistorialAsignacionCamion, error)
	GetActivoByCamionId(camionId int32) (*entities.HistorialAsignacionCamion, error)
	GetActivoByChoferId(choferId int32) (*entities.HistorialAsignacionCamion, error)
	DarDeBaja(idHistorial int32) error
	CerrarAsignacionActivaCamion(camionId int32) error
	CerrarAsignacionActivaChofer(choferId int32) error
}
