package infrastructure

import (
	"github.com/gin-gonic/gin"
	controllers "github.com/vicpoo/API_recolecta/src/usuario/infrastructure/controller"
	"go.uber.org/dig"
)


type UsuarioRoutesParams struct {
	dig.In

	Router *gin.Engine

	Create *controllers.AddUsersController
	Get    *controllers.ViewOneUsersController
	List   *controllers.ViewAllUsersController
	Delete *controllers.DeleteUsersController
	Login  *controllers.LoginUsersController
}

type UsuarioRoutes struct {
	router *gin.Engine

	create *controllers.AddUsersController
	get    *controllers.ViewOneUsersController
	list   *controllers.ViewAllUsersController
	delete *controllers.DeleteUsersController
	login  *controllers.LoginUsersController
}


func NewUsuarioRoutes(params UsuarioRoutesParams) *UsuarioRoutes {
	return &UsuarioRoutes{
		router: params.Router,
		create: params.Create,
		get:    params.Get,
		list:   params.List,
		delete: params.Delete,
		login:  params.Login,
	}
}

func (ur *UsuarioRoutes) Register() {
	usuariosGroup := ur.router.Group("/api/usuarios")

	usuariosGroup.POST("", ur.create.Handle)
	usuariosGroup.GET("/:id", ur.get.Handle)
	usuariosGroup.GET("", ur.list.Handle)
	usuariosGroup.DELETE("/:id", ur.delete.Handle)
	usuariosGroup.POST("/login", ur.login.Handle)
}