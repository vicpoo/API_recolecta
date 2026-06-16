// seguimiento_falla_critica_routes.go
package infrastructure

import (
	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/core"
)

type SeguimientoFallaCriticaRouter struct {
	engine *gin.Engine
}

func NewSeguimientoFallaCriticaRouter(engine *gin.Engine) *SeguimientoFallaCriticaRouter {
	return &SeguimientoFallaCriticaRouter{
		engine: engine,
	}
}

func (router *SeguimientoFallaCriticaRouter) Run() {
	// Inicializar dependencias
	createController, getByIdController, updateController, deleteController,
		getAllController, getByFallaIDController, getByFechaRangeController := InitSeguimientoFallaCriticaDependencies()

	// Grupo de rutas para seguimientos de falla crítica con prefijo /api
	seguimientoFallaCriticaGroup := router.engine.Group("/api/seguimientos-falla-critica")
	seguimientoFallaCriticaGroup.Use(core.JWTAuthMiddleware(), core.RequireRole(core.ADMIN, core.SUPERVISOR, core.COORDINADOR))
	{
		// Rutas CRUD básicas
		seguimientoFallaCriticaGroup.POST("/", createController.Run)
		seguimientoFallaCriticaGroup.GET("/:id", getByIdController.Run)
		seguimientoFallaCriticaGroup.PUT("/:id", updateController.Run)
		seguimientoFallaCriticaGroup.DELETE("/:id", deleteController.Run)
		seguimientoFallaCriticaGroup.GET("/", getAllController.Run)

		// Rutas específicas
		seguimientoFallaCriticaGroup.GET("/falla/:fallaId", getByFallaIDController.Run)
		seguimientoFallaCriticaGroup.GET("/por-fecha", getByFechaRangeController.Run)
	}
}
