package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/EstadoCamion/application"
)

type GetEstadoCamionByIdController struct {
	uc *application.GetByIdEstadoCamionUseCase
}

func NewGetEstadoCamionByIdController(
	uc *application.GetByIdEstadoCamionUseCase,
) *GetEstadoCamionByIdController {
	return &GetEstadoCamionByIdController{
		uc: uc,
	}
}

func (ctr *GetEstadoCamionByIdController) Run(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil || id <= 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "id inválido",
		})
		return
	}

	estadoCamion, err := ctr.uc.Run(int32(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    estadoCamion,
	})
}
