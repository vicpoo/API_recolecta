package domain

type DomicilioRepository interface {
	Create(d *Domicilio) error
	GetByID(id int) (*Domicilio, error)
	Update(d *Domicilio) error
	Delete(id int, usuarioID int) error
}
