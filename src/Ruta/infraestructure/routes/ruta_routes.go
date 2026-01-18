package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/Ruta/infraestructure/controllers"
)

type RutaRoutes struct {
	engine *gin.Engine

	createController  *controllers.CreateRutaController
	getAllController  *controllers.GetAllRutaController
	getByIdController *controllers.GetRutaByIdController
	updateController  *controllers.UpdateRutaController
	deleteController  *controllers.DeleteRutaController
}

func NewRutaRoutes(
	engine *gin.Engine,
	createController *controllers.CreateRutaController,
	getAllController *controllers.GetAllRutaController,
	getByIdController *controllers.GetRutaByIdController,
	updateController *controllers.UpdateRutaController,
	deleteController *controllers.DeleteRutaController,
) *RutaRoutes {
	return &RutaRoutes{
		engine: engine,

		createController:  createController,
		getAllController:  getAllController,
		getByIdController: getByIdController,
		updateController:  updateController,
		deleteController:  deleteController,
	}
}

func (r *RutaRoutes) Run() {
	routes := r.engine.Group("/api/rutas")
	{
		routes.POST("/", r.createController.Run)
		routes.GET("/", r.getAllController.Run)
		routes.GET("/:id", r.getByIdController.Run)
		routes.PUT("/:id", r.updateController.Run)
		routes.DELETE("/:id", r.deleteController.Run)
	}
}
