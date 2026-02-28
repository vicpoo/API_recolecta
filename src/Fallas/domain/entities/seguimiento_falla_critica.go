// seguimiento_falla_critica.go
package entities

import (
	"time"
)

type SeguimientoFallaCritica struct {
	SeguimientoID int32     `json:"seguimiento_id" gorm:"column:seguimiento_id;primaryKey;autoIncrement"`
	FallaID       int32     `json:"falla_id" gorm:"column:falla_id;not null"`
	Comentario    string    `json:"comentario" gorm:"column:comentario;type:text;not null"`
	CreatedAt     time.Time `json:"created_at" gorm:"column:created_at"`
}

// Setters
func (s *SeguimientoFallaCritica) SetSeguimientoID(seguimientoID int32) {
	s.SeguimientoID = seguimientoID
}

func (s *SeguimientoFallaCritica) SetFallaID(fallaID int32) {
	s.FallaID = fallaID
}

func (s *SeguimientoFallaCritica) SetComentario(comentario string) {
	s.Comentario = comentario
}

func (s *SeguimientoFallaCritica) SetCreatedAt(createdAt time.Time) {
	s.CreatedAt = createdAt
}

// Getters
func (s *SeguimientoFallaCritica) GetSeguimientoID() int32 {
	return s.SeguimientoID
}

func (s *SeguimientoFallaCritica) GetFallaID() int32 {
	return s.FallaID
}

func (s *SeguimientoFallaCritica) GetComentario() string {
	return s.Comentario
}

func (s *SeguimientoFallaCritica) GetCreatedAt() time.Time {
	return s.CreatedAt
}

// Constructor básico
func NewSeguimientoFallaCritica(fallaID int32, comentario string) *SeguimientoFallaCritica {
	return &SeguimientoFallaCritica{
		FallaID:    fallaID,
		Comentario: comentario,
		CreatedAt:  time.Now(),
	}
}

// Constructor completo
func NewSeguimientoFallaCriticaCompleto(seguimientoID int32, fallaID int32, comentario string, createdAt time.Time) *SeguimientoFallaCritica {
	return &SeguimientoFallaCritica{
		SeguimientoID: seguimientoID,
		FallaID:       fallaID,
		Comentario:    comentario,
		CreatedAt:     createdAt,
	}
}

// Método para agregar más comentarios
func (s *SeguimientoFallaCritica) AgregarComentario(comentarioAdicional string) {
	if s.Comentario == "" {
		s.Comentario = comentarioAdicional
	} else {
		s.Comentario = s.Comentario + "\n\n" + comentarioAdicional
	}
}

// Método para obtener fecha formateada
func (s *SeguimientoFallaCritica) GetFechaFormateada() string {
	return s.CreatedAt.Format("02/01/2006 15:04:05")
}

// Método para verificar si el comentario está vacío
func (s *SeguimientoFallaCritica) TieneComentario() bool {
	return s.Comentario != ""
}

// Método para obtener una vista previa del comentario (primeros 100 caracteres)
func (s *SeguimientoFallaCritica) GetComentarioResumido() string {
	if len(s.Comentario) <= 100 {
		return s.Comentario
	}
	return s.Comentario[:100] + "..."
}

// TableName especifica el nombre de la tabla para GORM
func (SeguimientoFallaCritica) TableName() string {
	return "seguimiento_falla_critica"
}