// incidencia_routes.go
package infrastructure

import (
	"github.com/gin-gonic/gin"
)

type IncidenciaRouter struct {
	engine *gin.Engine
}

func NewIncidenciaRouter(engine *gin.Engine) *IncidenciaRouter {
	return &IncidenciaRouter{
		engine: engine,
	}
}

func (router *IncidenciaRouter) Run() {
	// Inicializar dependencias
	createController, getByIdController, updateController, deleteController, getAllController,
	getByConductorController, getByPuntoController, getByFechaController :=
	InitIncidenciaDependencies()

	// Grupo de rutas para incidencias con prefijo /api
	incidenciaGroup := router.engine.Group("/api/incidencias")
	{
		// Rutas CRUD básicas
		incidenciaGroup.POST("/", createController.Run)
		incidenciaGroup.GET("/", getAllController.Run)
		incidenciaGroup.GET("/:id", getByIdController.Run)
		incidenciaGroup.PUT("/:id", updateController.Run)
		incidenciaGroup.DELETE("/:id", deleteController.Run)
		
		// Rutas específicas por filtros
		incidenciaGroup.GET("/conductor/:conductor_id", getByConductorController.Run)
		incidenciaGroup.GET("/punto/:punto_recoleccion_id", getByPuntoController.Run)
		incidenciaGroup.GET("/fecha", getByFechaController.Run)
	}
}