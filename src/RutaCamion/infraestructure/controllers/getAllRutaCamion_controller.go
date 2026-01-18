package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/RutaCamion/application"
)

type GetAllRutaCamionController struct {
	uc *application.ListAllRutaCamionUseCase
}

func NewGetAllRutaCamionController(
	uc *application.ListAllRutaCamionUseCase,
) *GetAllRutaCamionController {
	return &GetAllRutaCamionController{uc}
}

func (c *GetAllRutaCamionController) Run(ctx *gin.Context) {
	result, err := c.uc.Execute()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, result)
}
