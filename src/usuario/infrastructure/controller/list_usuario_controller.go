package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/usuario/application"
)

type ViewAllUsersController struct {
	uc *application.ViewAllUser
}

func NewViewAllUsersController(uc *application.ViewAllUser) *ViewAllUsersController {
	return &ViewAllUsersController{uc: uc}
}

// @Summary      Listar usuarios
// @Tags         Usuario
// @Produce      json
// @Success      200 {array} map[string]interface{}
// @Router       /api/usuarios [get]
func (c *ViewAllUsersController) Handle(ctx *gin.Context) {
	usuarios, err := c.uc.Execute(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, usuarios)
}
