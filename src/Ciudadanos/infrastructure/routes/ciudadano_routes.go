package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/core"
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
	ciudadanos := router.Group("/api/ciudadanos")

	ciudadanos.POST("", createController.Run)
	ciudadanos.POST("/login", loginController.Run)

	protected := ciudadanos.Group("")
	protected.Use(core.JWTAuthMiddleware())
	protected.Use(core.RequireRole(core.ADMIN))
	{
		protected.GET("", listController.Run)
		protected.GET("/:id", getController.Run)
		protected.PATCH("/:id", updateController.Run)
		protected.DELETE("/:id", deleteController.Run)
	}
}