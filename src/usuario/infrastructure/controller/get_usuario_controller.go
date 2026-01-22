package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/usuario/application"
)

type ViewOneUsersController struct {
	uc *application.ViewOneUser
}

func NewViewOneUsersController(uc *application.ViewOneUser) *ViewOneUsersController {
	return &ViewOneUsersController{uc: uc}
}

func (c *ViewOneUsersController) Handle(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil || id <= 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "id inválido"})
		return
	}

	usuario, err := c.uc.Execute(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "usuario no encontrado"})
		return
	}

	ctx.JSON(http.StatusOK, usuario)
}
