package domain

import "time"

type AlertaUsuario struct {
	AlertaID  int       `json:"alerta_id"`
	UsuarioID int       `json:"usuario_id"`   // receptor
	Titulo    string    `json:"titulo"`
	Mensaje   string    `json:"mensaje"`
	Leida     bool      `json:"leida"`
	CreadoPor int       `json:"creado_por"`   // supervisor
	CreatedAt time.Time `json:"created_at"`
}
