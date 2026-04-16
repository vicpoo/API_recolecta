package infrastructure

import "github.com/gin-gonic/gin"

func RegisterCiudadanoRoutes(engine *gin.Engine, deps *CiudadanoDependencies) {
	deps.Controller.RegisterRoutes(engine)
}
