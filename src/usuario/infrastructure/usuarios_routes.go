package infrastructure

import "github.com/gin-gonic/gin"

func RegisterUsuarioRoutes(r *gin.Engine, deps *UsuarioDependencies) {
	usuariosGroup := r.Group("/api/usuarios")
	
	usuariosGroup.POST("", deps.Create.Handle)
	usuariosGroup.GET("/:id", deps.Get.Handle)
	usuariosGroup.GET("", deps.List.Handle)
	usuariosGroup.DELETE("/:id", deps.Delete.Handle)
	usuariosGroup.POST("/login", deps.Login.Handle)
}
