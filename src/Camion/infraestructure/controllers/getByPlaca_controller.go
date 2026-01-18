package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/Camion/application"
)

type GetCamionByPlacaController struct {
	uc *application.GetCamionByPlacaUseCase
}

func NewGetCamionByPlacaController(uc *application.GetCamionByPlacaUseCase) *GetCamionByPlacaController {
	return &GetCamionByPlacaController{uc: uc}
}

func (ctr *GetCamionByPlacaController) Run(ctx *gin.Context) {
	placa := ctx.Param("placa")

	camion, err := ctr.uc.Run(placa)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    camion,
	})
}
