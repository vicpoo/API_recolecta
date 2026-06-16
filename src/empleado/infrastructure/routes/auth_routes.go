package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/vicpoo/API_recolecta/src/core"
	"github.com/vicpoo/API_recolecta/src/empleado/infrastructure/controller"
)

func EmpleadoRoutes(
	router *gin.Engine,
	createController *controller.CreateEmpleadoController,
	listController *controller.ListEmpleadoController,
	getController *controller.GetEmpleadoController,
	updateController *controller.UpdateEmpleadoController,
	deleteController *controller.DeleteEmpleadoController,
	loginController *controller.LoginEmpleadoController,
) {
	empleados := router.Group("/api/empleados")
	empleados.POST("/login", loginController.Run)

	protected := empleados.Group("")
	protected.Use(core.JWTAuthMiddleware(), core.RequireRole(core.ADMIN))
	{
		protected.POST("/", createController.Run)
		protected.GET("/", listController.Run)
		protected.GET("/:id", getController.Run)
		protected.PATCH("/:id", updateController.Run)
		protected.DELETE("/:id", deleteController.Run)
	}
}
