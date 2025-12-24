
package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/Camion/application"
)

type GetAllCamionController struct {
	uc *application.ListCamionUseCase
}

func NewGetAllCamionController(uc *application.ListCamionUseCase) *GetAllCamionController {
	return &GetAllCamionController{
		uc: uc,
	}
}

func (ctr *GetAllCamionController) Run(ctx *gin.Context) {
	camiones, err := ctr.uc.Run()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "error al obtener los camiones",
			"error":   err.Error(),
			"code":    http.StatusInternalServerError,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "lista de camiones obtenida correctamente",
		"data":    camiones,
		"code":    http.StatusOK,
	})
}
