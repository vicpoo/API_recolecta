package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/TipoCamion/infraestructure/controllers"
)


type TipoCamionRoutes struct {
	engine *gin.Engine
	createTipoCamionController *controllers.CreateTipoCamionController
	getTipoCamionController *controllers.GetAllTipoCamionController
	getTipoCamionByNameController *controllers.GetTipoCamionByNameController
	deleteTipoCamionById *controllers.DeleteTipoCamionController
}

func NewTipoCamionRoutes(engine *gin.Engine, createTipoCamionController *controllers.CreateTipoCamionController, getTipoCamionController *controllers.GetAllTipoCamionController, getTipoCamionByNameController *controllers.GetTipoCamionByNameController, deleteTipoCamionById *controllers.DeleteTipoCamionController) *TipoCamionRoutes {
	return &TipoCamionRoutes{
		engine: engine,
		createTipoCamionController: createTipoCamionController,
		getTipoCamionController: getTipoCamionController,
		getTipoCamionByNameController: getTipoCamionByNameController,
		deleteTipoCamionById: deleteTipoCamionById,
	}
}

func (tipoCamionRoutes *TipoCamionRoutes) Run() {
	routes := tipoCamionRoutes.engine.Group("/tipo-camion")
	{
		routes.POST("/", tipoCamionRoutes.createTipoCamionController.Run)
		routes.GET("/", tipoCamionRoutes.getTipoCamionController.Run)
		routes.GET("/nombre/:nombre", tipoCamionRoutes.getTipoCamionByNameController.Run)
		routes.DELETE("/:id", tipoCamionRoutes.deleteTipoCamionById.Run)
	}
}