package ports

import "github.com/vicpoo/API_recolecta/src/Ruta/domain/entities"

type IRuta interface {
	Save(ruta *entities.Ruta) error
	ListAll() ([]entities.Ruta, error)
	GetById(id int32) (*entities.Ruta, error)
	Update(ruta *entities.Ruta) error
	Delete(id int32) error
	GetActivas() ([]entities.Ruta, error)
}
