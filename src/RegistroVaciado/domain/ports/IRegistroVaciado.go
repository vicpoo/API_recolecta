package ports

import "github.com/vicpoo/API_recolecta/src/RegistroVaciado/domain/entities"

type RegistroVaciadoRepository interface {
	Save(registro *entities.RegistroVaciado) (*entities.RegistroVaciado, error)

	// READ
	ListAll() ([]entities.RegistroVaciado, error)
	GetByID(id int32) (*entities.RegistroVaciado, error)
	GetByRellenoID(rellenoID int32) ([]entities.RegistroVaciado, error)
	GetByRutaCamionID(rutaCamionID int32) ([]entities.RegistroVaciado, error)

	// DELETE
	Delete(id int32) error

	// UTILS
	ExistsByID(id int32) (bool, error)
}
