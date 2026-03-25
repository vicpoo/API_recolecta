package entities

import "time"

type CiudadanoPostgres struct {
	ID           int       `json:"id"`
	Email        string    `json:"email"`
	Alias        string    `json:"alias"`
	PasswordHash string    `json:"-"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
