package infrastructure

import (
	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/rol/infrastructure/controller"
	"go.uber.org/dig"
)

//////////////////////////////////////
// DEPENDENCIAS PARA INYECCIÓN
//////////////////////////////////////

type RolRoutesParams struct {
	dig.In

	Router        *gin.Engine
	RolController *controller.RolController
}

//////////////////////////////////////
// CLASE ROUTES
//////////////////////////////////////

type RolRoutes struct {
	router        *gin.Engine
	rolController *controller.RolController
}

//////////////////////////////////////
// CONSTRUCTOR PARA DI
//////////////////////////////////////

func NewRolRoutes(params RolRoutesParams) *RolRoutes {
	return &RolRoutes{
		router:        params.Router,
		rolController: params.RolController,
	}
}

//////////////////////////////////////
// MÉTODO PARA REGISTRAR RUTAS
//////////////////////////////////////

func (rr *RolRoutes) Register() {
	rolGroup := rr.router.Group("/api")
	rr.rolController.RegisterRoutes(rolGroup)
}