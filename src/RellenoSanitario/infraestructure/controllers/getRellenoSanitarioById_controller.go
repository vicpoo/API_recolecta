package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/RellenoSanitario/application"
)

type GetRellenoSanitarioByIDController struct {
	uc *application.GetRellenoSanitarioByIdUseCase
}

func NewGetRellenoSanitarioByIDController(uc *application.GetRellenoSanitarioByIdUseCase) *GetRellenoSanitarioByIDController {
	return &GetRellenoSanitarioByIDController{uc: uc}
}

func (c *GetRellenoSanitarioByIDController) Execute(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	result, err := c.uc.Execute(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, result)
}
