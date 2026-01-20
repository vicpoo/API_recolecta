package domain

import "time"

type Usuario struct {
	UserID       int        `json:"user_id"`
	Nombre       string     `json:"nombre"`
	Alias        string     `json:"alias"`
	Telefono     string     `json:"telefono"`
	Email        string     `json:"email"`
	Password     string     `json:"password"`
	RoleID       int        `json:"role_id"`
	ResidenciaID *int       `json:"residencia_id,omitempty"`
	Eliminado    bool       `json:"eliminado"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`
}
