package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/TipoCamion/application"
)

type GetTipoCamionByNameController struct {
	uc *application.GetTipoCamionByNameUseCase
}

func NewGetTipoCamionByNameController(
	uc *application.GetTipoCamionByNameUseCase,
) *GetTipoCamionByNameController {
	return &GetTipoCamionByNameController{uc: uc}
}

func (ctr *GetTipoCamionByNameController) Run(ctx *gin.Context) {
	nombre := ctx.Param("nombre")

	tipoCamion, err := ctr.uc.Run(nombre)

	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    tipoCamion,
	})
}
