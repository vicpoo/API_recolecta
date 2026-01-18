package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/RutaCamion/application"
)

type DeleteRutaCamionController struct {
	uc *application.DeleteRutaCamionUseCase
}

func NewDeleteRutaCamionController(
	uc *application.DeleteRutaCamionUseCase,
) *DeleteRutaCamionController {
	return &DeleteRutaCamionController{uc}
}

func (c *DeleteRutaCamionController) Run(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	if err := c.uc.Execute(int32(id)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "RutaCamion eliminada correctamente"})
}
