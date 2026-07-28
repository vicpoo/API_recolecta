package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/core"
	httpController "github.com/vicpoo/API_recolecta/src/Ciudadanos/infrastructure/controller/controller_domicilio"
)

func DomicilioRoutes(
	router *gin.Engine,
	domicilioController *httpController.DomicilioController,
) {
	
	domicilios := router.Group("/api/domicilios")
	domicilios.Use(core.JWTAuthMiddleware())

	{
		domicilios.POST("", domicilioController.Create)
		domicilios.GET("", domicilioController.List)
		domicilios.GET("/:id", domicilioController.GetByID)
		domicilios.PUT("/:id", domicilioController.Update)
		domicilios.DELETE("/:id", domicilioController.Delete)
	}
}