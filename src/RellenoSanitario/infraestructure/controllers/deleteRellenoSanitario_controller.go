package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/RellenoSanitario/application"
)

type DeleteRellenoSanitarioController struct {
	uc *application.DeleteRellenoSanitarioUseCase
}

func NewDeleteRellenoSanitarioController(uc *application.DeleteRellenoSanitarioUseCase) *DeleteRellenoSanitarioController {
	return &DeleteRellenoSanitarioController{uc: uc}
}

func (c *DeleteRellenoSanitarioController) Execute(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	if err := c.uc.Execute(id); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Relleno sanitario eliminado"})
}
