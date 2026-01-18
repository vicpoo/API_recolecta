package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/PuntoRecoleccion/application"
)

type DeletePuntoRecoleccionController struct {
	uc *application.DeletePuntoRecoleccionUseCase
}

func NewDeletePuntoRecoleccionController(uc *application.DeletePuntoRecoleccionUseCase) *DeletePuntoRecoleccionController {
	return &DeletePuntoRecoleccionController{uc: uc}
}

func (c *DeletePuntoRecoleccionController) Run(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "id inválido"})
		return
	}

	if err := c.uc.Execute(int32(id)); err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "punto de recolección eliminado"})
}
