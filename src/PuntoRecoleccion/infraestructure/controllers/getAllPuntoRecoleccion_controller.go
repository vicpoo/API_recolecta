package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/PuntoRecoleccion/application"
)

type GetAllPuntoRecoleccionController struct {
	uc *application.ListAllPuntoRecoleccionUseCase
}

func NewGetAllPuntoRecoleccionController(uc *application.ListAllPuntoRecoleccionUseCase) *GetAllPuntoRecoleccionController {
	return &GetAllPuntoRecoleccionController{uc: uc}
}

func (c *GetAllPuntoRecoleccionController) Run(ctx *gin.Context) {
	result, err := c.uc.Execute()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, result)
}
