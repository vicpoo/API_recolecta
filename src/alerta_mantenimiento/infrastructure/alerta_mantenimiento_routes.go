// alerta_mantenimiento_routes.go
package infrastructure

import (
	"github.com/gin-gonic/gin"
)

type AlertaMantenimientoRouter struct {
	engine *gin.Engine
}

func NewAlertaMantenimientoRouter(engine *gin.Engine) *AlertaMantenimientoRouter {
	return &AlertaMantenimientoRouter{
		engine: engine,
	}
}

func (router *AlertaMantenimientoRouter) Run() {
	// Inicializar dependencias
	createCtrl, getByIdCtrl, updateCtrl, deleteCtrl, getAllCtrl, getByCamionCtrl, getByTipoCtrl, getPendientesCtrl, getAtendidasCtrl, marcarAtendidaCtrl, getByFechaCtrl := InitAlertaMantenimientoDependencies()

	// Grupo de rutas para alertas de mantenimiento con prefijo /api
	alertaGroup := router.engine.Group("/api/alertas-mantenimiento")
	{
		// CRUD básico
		alertaGroup.POST("/", createCtrl.Run)
		alertaGroup.GET("/", getAllCtrl.Run)
		alertaGroup.GET("/:id", getByIdCtrl.Run)
		alertaGroup.PUT("/:id", updateCtrl.Run)
		alertaGroup.DELETE("/:id", deleteCtrl.Run)
		
		// Rutas específicas por estado
		alertaGroup.GET("/pendientes", getPendientesCtrl.Run)
		alertaGroup.GET("/atendidas", getAtendidasCtrl.Run)
		
		// Acciones específicas
		alertaGroup.PATCH("/:id/atender", marcarAtendidaCtrl.Run)
		
		// Rutas de filtrado
		alertaGroup.GET("/camion/:camion_id", getByCamionCtrl.Run)
		alertaGroup.GET("/tipo/:tipo_id", getByTipoCtrl.Run)
		alertaGroup.GET("/fecha", getByFechaCtrl.Run)
	}
}