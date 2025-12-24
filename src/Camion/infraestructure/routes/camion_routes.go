package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/Camion/infraestructure/controllers"
)

type CamionRoutes struct {
	engine *gin.Engine

	createCamionController   *controllers.CreateCamionController
	getAllCamionController   *controllers.GetAllCamionController
	getCamionByIdController  *controllers.GetCamionByIDController
	deleteCamionController   *controllers.DeleteCamionController
	getCamionByPlaca         *controllers.GetCamionByPlacaController
	getCamionByModelo        *controllers.GetCamionByModeloController
}

func NewCamionRoutes(
	engine *gin.Engine,
	createCamionController *controllers.CreateCamionController,
	getAllCamionController *controllers.GetAllCamionController,
	getCamionByIdController *controllers.GetCamionByIDController,
	deleteCamionController *controllers.DeleteCamionController,
	getCamionByPlaca       *controllers.GetCamionByPlacaController,
	getCamionByModelo      *controllers.GetCamionByModeloController, 
) *CamionRoutes {
	return &CamionRoutes{
		engine: engine,

		createCamionController:  createCamionController,
		getAllCamionController:  getAllCamionController,
		getCamionByIdController: getCamionByIdController,
		deleteCamionController:  deleteCamionController,
	}
}

func (camionRoutes *CamionRoutes) Run() {
	routes := camionRoutes.engine.Group("/camion")
	{
		routes.POST("/", camionRoutes.createCamionController.Run)
		routes.GET("/", camionRoutes.getAllCamionController.Run)
		routes.GET("/:id", camionRoutes.getCamionByIdController.Run)
		routes.DELETE("/:id", camionRoutes.deleteCamionController.Run)
		routes.GET("/placa/:placa", camionRoutes.getCamionByPlaca.Run)
		routes.GET("/modelo", camionRoutes.getCamionByModelo.Run)
	}
}
