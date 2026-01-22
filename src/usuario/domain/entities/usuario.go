package entities

import "time"

type Usuario struct {
	ID           int       `json:"id"`
	Nombre       string    `json:"nombre"`
	Email        string    `json:"email"`
	PasswordHash string    `json:"-"` 
	RolID        int       `json:"rol_id"`
	CreatedAt    time.Time `json:"created_at"`
	Alias        *string   `json:"alias"`
	Telefono     *string   `json:"telefono"`
	ResidenciaID *int      `json:"residencia_id"`
	Eliminado    bool      `json:"eliminado"`
	UpdatedAt    time.Time `json:"updated_at"`
}
