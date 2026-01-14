package domain

type NotificacionRepository interface {
	Create(n *Notificacion) error
	GetByID(id int) (*Notificacion, error)
	GetByUsuario(usuarioID int) ([]Notificacion, error)
	Deactivate(id int) error
}
