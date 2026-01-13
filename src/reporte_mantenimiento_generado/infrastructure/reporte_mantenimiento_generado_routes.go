// reporte_mantenimiento_generado_routes.go
package infrastructure

import (
	"github.com/gin-gonic/gin"
)

type ReporteMantenimientoGeneradoRouter struct {
	engine *gin.Engine
}

func NewReporteMantenimientoGeneradoRouter(engine *gin.Engine) *ReporteMantenimientoGeneradoRouter {
	return &ReporteMantenimientoGeneradoRouter{
		engine: engine,
	}
}

func (router *ReporteMantenimientoGeneradoRouter) Run() {
	// Inicializar dependencias
	createController, getByIdController, updateController, deleteController, getAllController,
	getByCoordinadorIDController, getByFechaRangeController, getByFechaGeneracionRangeController := InitReporteMantenimientoGeneradoDependencies()

	// Grupo de rutas para reportes de mantenimiento generados con prefijo /api
	reporteMantenimientoGeneradoGroup := router.engine.Group("/api/reportes-mantenimiento-generado")
	{
		// Operaciones CRUD básicas
		reporteMantenimientoGeneradoGroup.POST("/", createController.Run)
		reporteMantenimientoGeneradoGroup.GET("/", getAllController.Run)
		reporteMantenimientoGeneradoGroup.GET("/:id", getByIdController.Run)
		reporteMantenimientoGeneradoGroup.PUT("/:id", updateController.Run)
		reporteMantenimientoGeneradoGroup.DELETE("/:id", deleteController.Run)

		// Rutas específicas para filtros
		reporteMantenimientoGeneradoGroup.GET("/coordinador/:coordinador_id", getByCoordinadorIDController.Run)
		reporteMantenimientoGeneradoGroup.GET("/fecha", getByFechaRangeController.Run) // Filtra por rango del reporte (fecha_desde - fecha_hasta)
		reporteMantenimientoGeneradoGroup.GET("/fecha-generacion", getByFechaGeneracionRangeController.Run) // Filtra por fecha de creación (created_at)
	}
}