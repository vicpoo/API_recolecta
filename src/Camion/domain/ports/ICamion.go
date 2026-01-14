package ports

import "github.com/vicpoo/API_recolecta/src/Camion/domain/entities"

type ICamion interface {
	Save(camion *entities.Camion) (*entities.Camion, error)
	ListAll() ([]entities.Camion, error)
	GetByID(id int32) (*entities.Camion, error)
	Delete(id int32) error
	Update(id int32, camion *entities.Camion) (*entities.Camion, error)
	GetByPlaca(placa string) (*entities.Camion, error)
	GetByModelo(modelo string) ([]entities.Camion, error)
}
