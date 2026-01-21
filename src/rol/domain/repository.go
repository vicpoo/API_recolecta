package domain

type RolRepository interface {
	Create(nombre string) error
	List() ([]Rol, error)
	Update(id int, nombre string) error
}
