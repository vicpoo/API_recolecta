package domain

import "time"

type Domicilio struct {
	DomicilioID int       `json:"domicilio_id"`
	UsuarioID   int       `json:"usuario_id"`
	Alias       string    `json:"alias"`
	Direccion   string    `json:"direccion"`
	ColoniaID   int       `json:"colonia_id"`
	Eliminado   bool      `json:"eliminado"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
