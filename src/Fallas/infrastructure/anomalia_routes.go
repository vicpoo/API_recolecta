// anomalia_routes.go
package infrastructure

import (
	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/core"
	alertaDomain "github.com/vicpoo/API_recolecta/src/alerta_usuario/domain"
)

type AnomaliaRouter struct {
	engine     *gin.Engine
	alertaRepo alertaDomain.AlertaUsuarioRepository
}

func NewAnomaliaRouter(engine *gin.Engine, alertaRepo alertaDomain.AlertaUsuarioRepository) *AnomaliaRouter {
	return &AnomaliaRouter{
		engine:     engine,
		alertaRepo: alertaRepo,
	}
}

func (router *AnomaliaRouter) Run() {
	// Inicializar dependencias
	createController, getByIdController, updateController, deleteController,
		getAllController, getByPuntoIDController, getByChoferIDController,
		getByCamionIDController, getByRutaIDController, getByReferenciaIDController,
		getByEstadoController, getByTipoAnomaliaController, getByFechaRangeController := InitAnomaliaDependencies(router.alertaRepo)

	// Grupo de rutas para anomalías con prefijo /api
	anomaliaGroup := router.engine.Group("/api/anomalias")
	anomaliaGroup.Use(core.JWTAuthMiddleware(), core.RequireRole(core.ADMIN, core.SUPERVISOR, core.COORDINADOR))
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
		anomaliaGroup.GET("/camion/:camionId", getByCamionIDController.Run)
		anomaliaGroup.GET("/ruta/:rutaId", getByRutaIDController.Run)
		anomaliaGroup.GET("/referencia/:referenciaId", getByReferenciaIDController.Run)
		anomaliaGroup.GET("/estado", getByEstadoController.Run)        // Query param: ?estado=
		anomaliaGroup.GET("/tipo", getByTipoAnomaliaController.Run)    // Query param: ?tipo_anomalia=
		anomaliaGroup.GET("/por-fecha", getByFechaRangeController.Run) // Query params: ?fecha_inicio=&fecha_fin=
	}
}
