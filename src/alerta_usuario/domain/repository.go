package domain

type AlertaUsuarioRepository interface {
	Create(a *AlertaUsuario) error
	GetByUsuario(usuarioID int) ([]AlertaUsuario, error)
	MarkAsRead(alertaID int, usuarioID int) error
}
