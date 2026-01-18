package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/RegistroVaciado/application"
)

type DeleteRegistroVaciadoController struct {
	uc *application.DeleteRegistroVaciadoUseCase
}

func NewDeleteRegistroVaciadoController(uc *application.DeleteRegistroVaciadoUseCase) *DeleteRegistroVaciadoController {
	return &DeleteRegistroVaciadoController{uc: uc}
}

func (c *DeleteRegistroVaciadoController) Run(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	if err := c.uc.Execute(int32(id)); err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Registro eliminado"})
}
