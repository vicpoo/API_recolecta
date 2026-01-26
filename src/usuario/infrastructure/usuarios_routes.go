package infrastructure

import "github.com/gin-gonic/gin"
import "github.com/vicpoo/API_recolecta/src/core"

func RegisterUsuarioRoutes(r *gin.Engine, deps *UsuarioDependencies) {
	usuariosGroup := r.Group("/api/usuarios")
	
	
	usuariosGroup.POST("", deps.Create.Handle)
	usuariosGroup.POST("/login", deps.Login.Handle)

	
	protected := usuariosGroup.Group("")
	protected.Use(core.JWTAuthMiddleware())
	{
		protected.GET("/:id", deps.Get.Handle)
		protected.GET("", deps.List.Handle)
		protected.PUT("/:id", deps.Update.Handle)
		protected.DELETE("/:id", deps.Delete.Handle)
	}
}
