package ports

import "github.com/vicpoo/API_recolecta/src/PuntoRecoleccion/domain/entities"

type IPuntoRecoleccion interface {
	Save(p *entities.PuntoRecoleccion) (*entities.PuntoRecoleccion, error)
	Update(id int32, p *entities.PuntoRecoleccion) (*entities.PuntoRecoleccion, error)
	ListAll() ([]entities.PuntoRecoleccion, error)
	GetById(id int32) (*entities.PuntoRecoleccion, error)
	GetByRuta(rutaId int32) ([]entities.PuntoRecoleccion, error)
	Delete(id int32) error
}
