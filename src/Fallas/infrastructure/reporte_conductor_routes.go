// reporte_conductor_routes.go
package infrastructure

import (
	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/core"
)

type ReporteConductorRouter struct {
	engine *gin.Engine
}

func NewReporteConductorRouter(engine *gin.Engine) *ReporteConductorRouter {
	return &ReporteConductorRouter{
		engine: engine,
	}
}

func (router *ReporteConductorRouter) Run() {
	// Inicializar dependencias
	createController, getByIdController, updateController, deleteController, getAllController,
		getByCamionIDController, getByConductorIDController, getByRutaIDController, getByFechaRangeController := InitReporteConductorDependencies()

	// Grupo de rutas para reportes de conductores con prefijo /api
	reporteConductorGroup := router.engine.Group("/api/reportes-conductor")
	reporteConductorGroup.Use(core.JWTAuthMiddleware(), core.RequireRole(core.CONDUCTOR, core.ADMIN, core.COORDINADOR, core.SUPERVISOR))
	{
		// Operaciones CRUD básicas
		reporteConductorGroup.POST("/", createController.Run)
		reporteConductorGroup.GET("/", getAllController.Run)
		reporteConductorGroup.GET("/:id", getByIdController.Run)
		reporteConductorGroup.PUT("/:id", updateController.Run)
		reporteConductorGroup.DELETE("/:id", deleteController.Run)

		// Rutas específicas para filtros
		reporteConductorGroup.GET("/camion/:camion_id", getByCamionIDController.Run)
		reporteConductorGroup.GET("/conductor/:conductor_id", getByConductorIDController.Run)
		reporteConductorGroup.GET("/ruta/:ruta_id", getByRutaIDController.Run)
		reporteConductorGroup.GET("/fecha", getByFechaRangeController.Run) // Usando query params
	}
}
