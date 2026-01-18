package domain

type UsuarioRepository interface {
	Create(usuario *Usuario) error
	GetByID(id int) (*Usuario, error)
	GetAll() ([]Usuario, error)
	GetByIDEmail(email string) (*Usuario, error)
	Delete(id int) error
}
