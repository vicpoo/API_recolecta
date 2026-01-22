package domain

import "github.com/vicpoo/API_recolecta/src/rol/domain/entities"

type RolRepository interface {
	Create(nombre string) error
	List() ([]entities.Rol, error)
	Update(id int, nombre string) error
	Delete(id int) error
}
