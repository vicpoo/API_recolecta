package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/Ruta/application"
)

type DeleteRutaController struct {
	uc *application.DeleteRutaUseCase
}

func NewDeleteRutaController(uc *application.DeleteRutaUseCase) *DeleteRutaController {
	return &DeleteRutaController{uc}
}

func (ctr *DeleteRutaController) Run(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	err := ctr.uc.Run(int32(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"success": false, "message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"success": true, "message": "ruta eliminada"})
}
