package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/RutaCamion/application"
)

type GetRutaCamionByIDController struct {
	uc *application.GetRutaCamionByIDUseCase
}

func NewGetRutaCamionByIDController(
	uc *application.GetRutaCamionByIDUseCase,
) *GetRutaCamionByIDController {
	return &GetRutaCamionByIDController{uc}
}

func (c *GetRutaCamionByIDController) Run(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	result, err := c.uc.Execute(int32(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, result)
}
