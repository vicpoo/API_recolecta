package domain

import "time"

type Domicilio struct {
	DomicilioID int
	UsuarioID   int
	Alias       string
	Direccion   string
	ColoniaID   int
	Eliminado   bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
