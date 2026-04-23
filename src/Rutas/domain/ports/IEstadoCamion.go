package ports

import "github.com/vicpoo/API_recolecta/src/Rutas/domain/entities"

type IEstadoCamion interface {
	Save(estado *entities.EstadoCamion) (*entities.EstadoCamion, error)
	ListAll() ([]entities.EstadoCamion, error)
	GetById(id int32) (*entities.EstadoCamion, error)
	Update(id int32,estado *entities.EstadoCamion) (*entities.EstadoCamion, error)
	Delete(id int32) error
}
