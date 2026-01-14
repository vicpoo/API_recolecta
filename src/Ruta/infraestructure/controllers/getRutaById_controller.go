package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/Ruta/application"
)

type GetRutaByIdController struct {
	uc *application.GetRutaByIdUseCase
}

func NewGetRutaByIdController(uc *application.GetRutaByIdUseCase) *GetRutaByIdController {
	return &GetRutaByIdController{
		uc: uc,
	}
}

func (ctr *GetRutaByIdController) Run(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "ID inválido",
		})
		return
	}

	ruta, err := ctr.uc.Run(int32(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"message": "Ruta no encontrada",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    ruta,
	})
}
