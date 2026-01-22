package infrastructure

import (
	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/rol/infrastructure/controller"
)

func RegisterRolRoutes(r *gin.Engine, rolController *controller.RolController) {
	rolGroup := r.Group("/api")
	rolController.RegisterRoutes(rolGroup)
}
