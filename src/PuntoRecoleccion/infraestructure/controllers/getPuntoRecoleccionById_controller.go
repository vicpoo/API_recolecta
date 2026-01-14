package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/PuntoRecoleccion/application"
)

type GetPuntoRecoleccionByIdController struct {
	uc *application.GetPuntoRecoleccionByIdUseCase
}

func NewGetPuntoRecoleccionByIdController(uc *application.GetPuntoRecoleccionByIdUseCase) *GetPuntoRecoleccionByIdController {
	return &GetPuntoRecoleccionByIdController{uc: uc}
}

func (c *GetPuntoRecoleccionByIdController) Run(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "id inválido"})
		return
	}

	result, err := c.uc.Execute(int32(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, result)
}
