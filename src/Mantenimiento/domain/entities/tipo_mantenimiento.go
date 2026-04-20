//tipo_mantenimiento.go
package entities

type TipoMantenimiento struct {
	ID        int32  `json:"tipo_mantenimiento_id" gorm:"column:tipo_mantenimiento_id;primaryKey;autoIncrement"`
	Nombre    string `json:"nombre" gorm:"column:nombre;not null"`
	Categoria string `json:"categoria" gorm:"column:categoria;not null"`
	Eliminado bool   `json:"eliminado" gorm:"column:eliminado"`
}

// Setters
func (t *TipoMantenimiento) SetID(id int32) {
	t.ID = id
}

func (t *TipoMantenimiento) SetNombre(nombre string) {
	t.Nombre = nombre
}

func (t *TipoMantenimiento) SetCategoria(categoria string) {
	t.Categoria = categoria
}

func (t *TipoMantenimiento) SetEliminado(eliminado bool) {
	t.Eliminado = eliminado
}

// Getters
func (t *TipoMantenimiento) GetID() int32 {
	return t.ID
}

func (t *TipoMantenimiento) GetNombre() string {
	return t.Nombre
}

func (t *TipoMantenimiento) GetCategoria() string {
	return t.Categoria
}

func (t *TipoMantenimiento) GetEliminado() bool {
	return t.Eliminado
}

// Constructor
func NewTipoMantenimiento(nombre string, categoria string) *TipoMantenimiento {
	return &TipoMantenimiento{
		Nombre:    nombre,
		Categoria: categoria,
		Eliminado: false,
	}
}


func (t *TipoMantenimiento) MarcarEliminado() {
	t.Eliminado = true
}