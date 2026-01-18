package domain

import "time"

type AlertaUsuario struct {
	AlertaID  int
	Titulo    string
	Mensaje   string
	UsuarioID int
	CreadoPor int
	Leida     bool
	CreatedAt time.Time
}
