package ports

import "github.com/vicpoo/API_recolecta/src/RutaCamion/domain/entities"

type RutaCamionRepository interface {
	Save(rutaCamion *entities.RutaCamion) (*entities.RutaCamion, error)
	Update(id int32, rutaCamion *entities.RutaCamion) (*entities.RutaCamion, error)
	ListAll() ([]entities.RutaCamion, error)
	GetByID(id int32) (*entities.RutaCamion, error)
	Delete(id int32) error
	GetByCamionID(camionID int32) ([]entities.RutaCamion, error)
	GetByRutaID(rutaID int32) ([]entities.RutaCamion, error)
	ExistsByID(id int32) (bool, error)
}
