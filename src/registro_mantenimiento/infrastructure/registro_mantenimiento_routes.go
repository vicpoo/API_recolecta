//registro_mantenimiento_routes.go
package infrastructure

import (
	"github.com/gin-gonic/gin"
)

type RegistroMantenimientoRouter struct {
	engine *gin.Engine
}

func NewRegistroMantenimientoRouter(engine *gin.Engine) *RegistroMantenimientoRouter {
	return &RegistroMantenimientoRouter{
		engine: engine,
	}
}

func (router *RegistroMantenimientoRouter) Run() {
	// Inicializar dependencias
	createController, getByIdController, updateController, deleteController, getAllController, 
	getByAlertaController, getByCamionController, getByCoordinadorController, getByFechaController := 
	InitRegistroMantenimientoDependencies()

	// Grupo de rutas para registros de mantenimiento con prefijo /api
	registroMantenimientoGroup := router.engine.Group("/api/registros-mantenimiento")
	{
		// Rutas CRUD básicas
		registroMantenimientoGroup.POST("/", createController.Run)
		registroMantenimientoGroup.GET("/", getAllController.Run)
		registroMantenimientoGroup.GET("/:id", getByIdController.Run)
		registroMantenimientoGroup.PUT("/:id", updateController.Run)
		registroMantenimientoGroup.DELETE("/:id", deleteController.Run)
		
		// Rutas específicas por filtros
		registroMantenimientoGroup.GET("/alerta/:alerta_id", getByAlertaController.Run)
		registroMantenimientoGroup.GET("/camion/:camion_id", getByCamionController.Run)
		registroMantenimientoGroup.GET("/coordinador/:coordinador_id", getByCoordinadorController.Run)
		registroMantenimientoGroup.GET("/fecha", getByFechaController.Run)
	}
}