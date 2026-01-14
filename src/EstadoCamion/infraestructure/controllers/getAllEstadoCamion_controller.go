package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/EstadoCamion/application"
)

type GetAllEstadoCamionController struct {
	uc *application.ListAllEstadoCamionUseCase
}

func NewGetAllEstadoCamionController(
	uc *application.ListAllEstadoCamionUseCase,
) *GetAllEstadoCamionController {
	return &GetAllEstadoCamionController{
		uc: uc,
	}
}

func (ctr *GetAllEstadoCamionController) Run(ctx *gin.Context) {
	estadosCamion, err := ctr.uc.Run()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "no se pudieron obtener los estados del camión",
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    estadosCamion,
	})
}
