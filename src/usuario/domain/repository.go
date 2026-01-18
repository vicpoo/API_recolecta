package domain

type UsuarioRepository interface {
	Create(usuario *Usuario) error
	GetByID(id int) (*Usuario, error)
	GetAll() ([]Usuario, error)
	GetByEmail(email string) (*Usuario, error)
	Delete(id int) error
}
