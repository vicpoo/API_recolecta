package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/Ciudadanos/infrastructure/controller/controller_ciudadano"
)

func CiudadanoRoutes(
	router *gin.Engine,
	createController *controller_ciudadano.CreateCiudadanoController,
	getController *controller_ciudadano.GetCiudadanoController,
	listController *controller_ciudadano.ListCiudadanoController,
	updateController *controller_ciudadano.UpdateCiudadanoController,
	deleteController *controller_ciudadano.DeleteCiudadanoController,
	loginController *controller_ciudadano.LoginCiudadanoController,
) {
	ciudadanos := router.Group("/ciudadanos")
	{
		ciudadanos.POST("", createController.Run)
		ciudadanos.GET("", listController.Run)
		ciudadanos.GET("/:id", getController.Run)
		ciudadanos.PUT("/:id", updateController.Run)
		ciudadanos.DELETE("/:id", deleteController.Run)
		ciudadanos.POST("/login", loginController.Run)
	}
}