package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/PuntoRecoleccion/application"
	"github.com/vicpoo/API_recolecta/src/PuntoRecoleccion/domain/entities"
)

type UpdatePuntoRecoleccionController struct {
	uc *application.UpdatePuntoRecoleccionUseCase
}

func NewUpdatePuntoRecoleccionController(uc *application.UpdatePuntoRecoleccionUseCase) *UpdatePuntoRecoleccionController {
	return &UpdatePuntoRecoleccionController{uc: uc}
}

func (c *UpdatePuntoRecoleccionController) Run(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "id inválido"})
		return
	}

	var p entities.PuntoRecoleccion
	if err := ctx.ShouldBindJSON(&p); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := c.uc.Execute(int32(id), &p)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, result)
}
