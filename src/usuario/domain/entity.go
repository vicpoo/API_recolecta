package domain

import "time"

type Usuario struct {
	UserID       int
	Nombre       string
	Alias        string
	Telefono     string
	Email        string
	Password     string
	RoleID       int
	ResidenciaID *int
	Eliminado    bool
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
