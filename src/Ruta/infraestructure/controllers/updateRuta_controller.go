package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/Ruta/domain/entities"
	"github.com/vicpoo/API_recolecta/src/Ruta/application"
)

type UpdateRutaController struct {
	uc *application.UpdateRutaUseCase
}

func NewUpdateRutaController(uc *application.UpdateRutaUseCase) *UpdateRutaController {
	return &UpdateRutaController{uc}
}

func (ctr *UpdateRutaController) Run(ctx *gin.Context) {
	var ruta entities.Ruta
	id, _ := strconv.Atoi(ctx.Param("id"))
	ruta.RutaID = int32(id)

	ctx.ShouldBindJSON(&ruta)

	err := ctr.uc.Run(&ruta)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"success": true, "message": "ruta actualizada"})
}
