package domain

import "time"

type Notificacion struct {
	NotificacionID          int
	UsuarioID               int
	Tipo                    string
	Titulo                  string
	Mensaje                 string
	Activa                  bool
	IDCamionRelacionado     *int
	IDFallaRelacionado      *int
	IDMantenimientoRelacionado *int
	CreadoPor               int
	CreatedAt               time.Time
}
