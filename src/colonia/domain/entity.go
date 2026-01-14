package domain

import "time"

type Colonia struct {
	ColoniaID int
	Nombre    string
	Zona      string
	CreatedAt time.Time
}
