// tipo_mantenimiento_routes.go
package infrastructure

import (
	"github.com/gin-gonic/gin"
)

type TipoMantenimientoRouter struct {
	engine *gin.Engine
}

func NewTipoMantenimientoRouter(engine *gin.Engine) *TipoMantenimientoRouter {
	return &TipoMantenimientoRouter{
		engine: engine,
	}
}

func (router *TipoMantenimientoRouter) Run() {
	// Inicializar dependencias
	createController, getByIdController, updateController, deleteController, getAllController := InitTipoMantenimientoDependencies()

	// Grupo de rutas para tipos de mantenimiento con prefijo /api
	tipoMantenimientoGroup := router.engine.Group("/api/tipos-mantenimiento")
	{
		tipoMantenimientoGroup.POST("/", createController.Run)
		tipoMantenimientoGroup.GET("/:id", getByIdController.Run)
		tipoMantenimientoGroup.PUT("/:id", updateController.Run)
		tipoMantenimientoGroup.DELETE("/:id", deleteController.Run)
		tipoMantenimientoGroup.GET("/", getAllController.Run)
	}
}