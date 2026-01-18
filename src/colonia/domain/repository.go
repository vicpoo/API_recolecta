package domain

type ColoniaRepository interface {
	Create(c *Colonia) error
	GetByID(id int) (*Colonia, error)
	GetAll() ([]Colonia, error)
	Update(c *Colonia) error
	Delete(id int) error
}
