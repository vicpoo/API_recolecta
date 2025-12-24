package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/Camion/application"
)

type GetCamionByModeloController struct {
	uc *application.GetCamionByModeloUseCase
}

func NewGetCamionByModeloController(uc *application.GetCamionByModeloUseCase) *GetCamionByModeloController {
	return &GetCamionByModeloController{uc: uc}
}

func (ctr *GetCamionByModeloController) Run(ctx *gin.Context) {
	modelo := ctx.Query("modelo")

	camiones, err := ctr.uc.Run(modelo)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    camiones,
	})
}
