// notificacion.go
package entities

import (
	"time"
	"errors"
)

type Notificacion struct {
	NotificacionID             int32     `json:"notificacion_id" gorm:"column:notificacion_id;primaryKey;autoIncrement"`
	UsuarioID                  *int32    `json:"usuario_id" gorm:"column:usuario_id"`
	Tipo                       string    `json:"tipo" gorm:"column:tipo;type:varchar(50);not null"`
	Titulo                     string    `json:"titulo" gorm:"column:titulo;type:varchar(100);not null"`
	Mensaje                    string    `json:"mensaje" gorm:"column:mensaje;type:text;not null"`
	Activa                     bool      `json:"activa" gorm:"column:activa;default:true"`
	IDCamionRelacionado        *int32    `json:"id_camion_relacionado,omitempty" gorm:"column:id_camion_relacionado"`
	IDFallaRelacionado         *int32    `json:"id_falla_relacionado,omitempty" gorm:"column:id_falla_relacionado"`
	IDMantenimientoRelacionado *int32    `json:"id_mantenimiento_relacionado,omitempty" gorm:"column:id_mantenimiento_relacionado"`
	CreadoPor                  *int32    `json:"creado_por" gorm:"column:creado_por"`
	CreatedAt                  time.Time `json:"created_at" gorm:"column:created_at"`
}

// TableName especifica el nombre de la tabla para GORM
func (Notificacion) TableName() string {
	return "notificacion"
}

// Setters
func (n *Notificacion) SetNotificacionID(notificacionID int32) {
	n.NotificacionID = notificacionID
}

func (n *Notificacion) SetUsuarioID(usuarioID *int32) {
	n.UsuarioID = usuarioID
}

func (n *Notificacion) SetTipo(tipo string) {
	n.Tipo = tipo
}

func (n *Notificacion) SetTitulo(titulo string) {
	n.Titulo = titulo
}

func (n *Notificacion) SetMensaje(mensaje string) {
	n.Mensaje = mensaje
}

func (n *Notificacion) SetActiva(activa bool) {
	n.Activa = activa
}

func (n *Notificacion) SetIDCamionRelacionado(idCamion *int32) {
	n.IDCamionRelacionado = idCamion
}

func (n *Notificacion) SetIDFallaRelacionado(idFalla *int32) {
	n.IDFallaRelacionado = idFalla
}

func (n *Notificacion) SetIDMantenimientoRelacionado(idMantenimiento *int32) {
	n.IDMantenimientoRelacionado = idMantenimiento
}

func (n *Notificacion) SetCreadoPor(creadoPor *int32) {
	n.CreadoPor = creadoPor
}

func (n *Notificacion) SetCreatedAt(createdAt time.Time) {
	n.CreatedAt = createdAt
}

// Getters
func (n *Notificacion) GetNotificacionID() int32 {
	return n.NotificacionID
}

func (n *Notificacion) GetUsuarioID() *int32 {
	return n.UsuarioID
}

func (n *Notificacion) GetTipo() string {
	return n.Tipo
}

func (n *Notificacion) GetTitulo() string {
	return n.Titulo
}

func (n *Notificacion) GetMensaje() string {
	return n.Mensaje
}

func (n *Notificacion) GetActiva() bool {
	return n.Activa
}

func (n *Notificacion) GetIDCamionRelacionado() *int32 {
	return n.IDCamionRelacionado
}

func (n *Notificacion) GetIDFallaRelacionado() *int32 {
	return n.IDFallaRelacionado
}

func (n *Notificacion) GetIDMantenimientoRelacionado() *int32 {
	return n.IDMantenimientoRelacionado
}

func (n *Notificacion) GetCreadoPor() *int32 {
	return n.CreadoPor
}

func (n *Notificacion) GetCreatedAt() time.Time {
	return n.CreatedAt
}

// Constructor básico
func NewNotificacion(tipo string, titulo string, mensaje string) *Notificacion {
	return &Notificacion{
		Tipo:      tipo,
		Titulo:    titulo,
		Mensaje:   mensaje,
		Activa:    true,
		CreatedAt: time.Now(),
	}
}

// Constructor para notificación genérica con usuario
func NewNotificacionUsuario(usuarioID int32, tipo string, titulo string, mensaje string, creadoPor *int32) *Notificacion {
	return &Notificacion{
		UsuarioID: &usuarioID,
		Tipo:      tipo,
		Titulo:    titulo,
		Mensaje:   mensaje,
		Activa:    true,
		CreadoPor: creadoPor,
		CreatedAt: time.Now(),
	}
}

// Constructor para notificación de falla (requiere camión y falla)
func NewNotificacionFalla(usuarioID *int32, titulo string, mensaje string, idCamion int32, idFalla int32, creadoPor *int32) *Notificacion {
	return &Notificacion{
		UsuarioID:           usuarioID,
		Tipo:                "falla",
		Titulo:              titulo,
		Mensaje:             mensaje,
		Activa:              true,
		IDCamionRelacionado: &idCamion,
		IDFallaRelacionado:  &idFalla,
		CreadoPor:           creadoPor,
		CreatedAt:           time.Now(),
	}
}

// Constructor para notificación de mantenimiento (requiere camión y mantenimiento)
func NewNotificacionMantenimiento(usuarioID *int32, titulo string, mensaje string, idCamion int32, idMantenimiento int32, creadoPor *int32) *Notificacion {
	return &Notificacion{
		UsuarioID:                  usuarioID,
		Tipo:                       "mantenimiento",
		Titulo:                     titulo,
		Mensaje:                    mensaje,
		Activa:                     true,
		IDCamionRelacionado:        &idCamion,
		IDMantenimientoRelacionado: &idMantenimiento,
		CreadoPor:                  creadoPor,
		CreatedAt:                  time.Now(),
	}
}

// Constructor para notificación de emergencia (solo requiere camión)
func NewNotificacionEmergencia(usuarioID *int32, titulo string, mensaje string, idCamion int32, creadoPor *int32) *Notificacion {
	return &Notificacion{
		UsuarioID:           usuarioID,
		Tipo:                "emergencia",
		Titulo:              titulo,
		Mensaje:             mensaje,
		Activa:              true,
		IDCamionRelacionado: &idCamion,
		CreadoPor:           creadoPor,
		CreatedAt:           time.Now(),
	}
}

// Constructor para notificación de ruta (puede requerir camión)
func NewNotificacionRuta(usuarioID *int32, titulo string, mensaje string, idCamion *int32, creadoPor *int32) *Notificacion {
	return &Notificacion{
		UsuarioID:           usuarioID,
		Tipo:                "ruta",
		Titulo:              titulo,
		Mensaje:             mensaje,
		Activa:              true,
		IDCamionRelacionado: idCamion,
		CreadoPor:           creadoPor,
		CreatedAt:           time.Now(),
	}
}

// Constructor completo
func NewNotificacionCompleta(notificacionID int32, usuarioID *int32, tipo string, titulo string, mensaje string, activa bool, idCamion *int32, idFalla *int32, idMantenimiento *int32, creadoPor *int32, createdAt time.Time) *Notificacion {
	return &Notificacion{
		NotificacionID:             notificacionID,
		UsuarioID:                  usuarioID,
		Tipo:                       tipo,
		Titulo:                     titulo,
		Mensaje:                    mensaje,
		Activa:                     activa,
		IDCamionRelacionado:        idCamion,
		IDFallaRelacionado:         idFalla,
		IDMantenimientoRelacionado: idMantenimiento,
		CreadoPor:                  creadoPor,
		CreatedAt:                  createdAt,
	}
}

// Método para verificar qué relaciones tiene
func (n *Notificacion) GetRelaciones() map[string]bool {
	return map[string]bool{
		"camion":        n.IDCamionRelacionado != nil,
		"falla":         n.IDFallaRelacionado != nil,
		"mantenimiento": n.IDMantenimientoRelacionado != nil,
	}
}

// Método para saber si tiene al menos una relación
func (n *Notificacion) TieneRelaciones() bool {
	return n.IDCamionRelacionado != nil || n.IDFallaRelacionado != nil || n.IDMantenimientoRelacionado != nil
}

// Método para limpiar todas las relaciones
func (n *Notificacion) LimpiarRelaciones() {
	n.IDCamionRelacionado = nil
	n.IDFallaRelacionado = nil
	n.IDMantenimientoRelacionado = nil
}

// Método para establecer solo relación con camión
func (n *Notificacion) SetRelacionCamion(idCamion int32) {
	n.LimpiarRelaciones()
	n.IDCamionRelacionado = &idCamion
}

// Método para establecer relación camión-falla
func (n *Notificacion) SetRelacionCamionFalla(idCamion int32, idFalla int32) {
	n.LimpiarRelaciones()
	n.IDCamionRelacionado = &idCamion
	n.IDFallaRelacionado = &idFalla
}

// Método para establecer relación camión-mantenimiento
func (n *Notificacion) SetRelacionCamionMantenimiento(idCamion int32, idMantenimiento int32) {
	n.LimpiarRelaciones()
	n.IDCamionRelacionado = &idCamion
	n.IDMantenimientoRelacionado = &idMantenimiento
}

// Método para marcar notificación como leída
func (n *Notificacion) MarcarComoLeida() {
	n.Activa = false
}

// Método para reactivar notificación
func (n *Notificacion) Reactivar() {
	n.Activa = true
}

// Método para actualizar contenido
func (n *Notificacion) ActualizarContenido(titulo string, mensaje string) {
	n.Titulo = titulo
	n.Mensaje = mensaje
}

// Método para obtener resumen
func (n *Notificacion) ObtenerResumen() string {
	if len(n.Mensaje) > 100 {
		return n.Mensaje[:100] + "..."
	}
	return n.Mensaje
}

// Método para hacer notificación global
func (n *Notificacion) HacerGlobal() {
	n.UsuarioID = nil
}

// Método para verificar si es global
func (n *Notificacion) EsGlobal() bool {
	return n.UsuarioID == nil
}

// Método para validar según el tipo
func (n *Notificacion) ValidarSegunTipo() error {
	switch n.Tipo {
	case "falla":
		if n.IDCamionRelacionado == nil || n.IDFallaRelacionado == nil {
			return errors.New("las notificaciones de falla requieren camión y falla relacionados")
		}
	case "mantenimiento":
		if n.IDCamionRelacionado == nil || n.IDMantenimientoRelacionado == nil {
			return errors.New("las notificaciones de mantenimiento requieren camión y mantenimiento relacionados")
		}
	case "emergencia":
		if n.IDCamionRelacionado == nil {
			return errors.New("las notificaciones de emergencia requieren camión relacionado")
		}
	// Para "ruta" y otros tipos, las relaciones son opcionales
	}
	return nil
}
