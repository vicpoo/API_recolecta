package ports

import "github.com/vicpoo/API_recolecta/src/Rutas/domain/entities"


type RellenoSanitarioRepository interface {
    Save(relleno *entities.RellenoSanitario) (*entities.RellenoSanitario, error)
    Update(id int32, relleno *entities.RellenoSanitario) (*entities.RellenoSanitario, error)
	ListAll() ([]entities.RellenoSanitario, error)
	GetByID(id int32) (*entities.RellenoSanitario, error)
	Delete(id int32) error
	GetByNombre(nombre string) ([]entities.RellenoSanitario, error)
	ExistsByID(id int32) (bool, error)
}