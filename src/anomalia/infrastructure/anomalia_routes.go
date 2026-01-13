// anomalia_routes.go
package infrastructure

import (
	"github.com/gin-gonic/gin"
)

type AnomaliaRouter struct {
	engine *gin.Engine
}

func NewAnomaliaRouter(engine *gin.Engine) *AnomaliaRouter {
	return &AnomaliaRouter{
		engine: engine,
	}
}

func (router *AnomaliaRouter) Run() {
	// Inicializar dependencias
	createController, getByIdController, updateController, deleteController,
	getAllController, getByPuntoIDController, getByChoferIDController,
	getByEstadoController, getByTipoAnomaliaController, getByFechaRangeController := InitAnomaliaDependencies()

	// Grupo de rutas para anomalías con prefijo /api
	anomaliaGroup := router.engine.Group("/api/anomalias")
	{
		// Rutas CRUD básicas
		anomaliaGroup.POST("/", createController.Run)
		anomaliaGroup.GET("/:id", getByIdController.Run)
		anomaliaGroup.PUT("/:id", updateController.Run)
		anomaliaGroup.DELETE("/:id", deleteController.Run)
		anomaliaGroup.GET("/", getAllController.Run)
		
		// Rutas específicas
		anomaliaGroup.GET("/punto/:puntoId", getByPuntoIDController.Run)
		anomaliaGroup.GET("/chofer/:choferId", getByChoferIDController.Run)
		anomaliaGroup.GET("/estado", getByEstadoController.Run) // Query param: ?estado=
		anomaliaGroup.GET("/tipo", getByTipoAnomaliaController.Run) // Query param: ?tipo_anomalia=
		anomaliaGroup.GET("/por-fecha", getByFechaRangeController.Run) // Query params: ?fecha_inicio=&fecha_fin=
	}
}