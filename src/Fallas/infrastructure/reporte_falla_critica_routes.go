// reporte_falla_critica_routes.go
package infrastructure

import (
	"github.com/gin-gonic/gin"
)

type ReporteFallaCriticaRouter struct {
	engine *gin.Engine
}

func NewReporteFallaCriticaRouter(engine *gin.Engine) *ReporteFallaCriticaRouter {
	return &ReporteFallaCriticaRouter{
		engine: engine,
	}
}

func (router *ReporteFallaCriticaRouter) Run() {
	// Inicializar dependencias
	createController, getByIdController, updateController, deleteController, 
	getAllController, getByCamionIDController, getByConductorIDController, 
	getByFechaRangeController := InitReporteFallaCriticaDependencies()

	// Grupo de rutas para reportes de falla crítica con prefijo /api
	reporteFallaCriticaGroup := router.engine.Group("/api/reportes-falla-critica")
	{
		// Rutas CRUD básicas
		reporteFallaCriticaGroup.POST("/", createController.Run)
		reporteFallaCriticaGroup.GET("/:id", getByIdController.Run)
		reporteFallaCriticaGroup.PUT("/:id", updateController.Run)
		reporteFallaCriticaGroup.DELETE("/:id", deleteController.Run)
		reporteFallaCriticaGroup.GET("/", getAllController.Run)
		
		// Rutas específicas
		reporteFallaCriticaGroup.GET("/camion/:camionId", getByCamionIDController.Run)
		reporteFallaCriticaGroup.GET("/conductor/:conductorId", getByConductorIDController.Run)
		reporteFallaCriticaGroup.GET("/por-fecha", getByFechaRangeController.Run)
	}
}