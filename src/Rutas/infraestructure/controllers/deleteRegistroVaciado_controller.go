package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/Rutas/application"
)

type DeleteRegistroVaciadoController struct {
	uc *application.DeleteRegistroVaciadoUseCase
}

func NewDeleteRegistroVaciadoController(uc *application.DeleteRegistroVaciadoUseCase) *DeleteRegistroVaciadoController {
	return &DeleteRegistroVaciadoController{uc: uc}
}

// @Summary      Eliminar registro de vaciado
// @Tags         RegistroVaciado
// @Produce      json
// @Param        id path int true "ID"
// @Success      200 {object} map[string]interface{}
// @Router       /api/registro-vaciado/{id} [delete]
func (c *DeleteRegistroVaciadoController) Run(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	if err := c.uc.Execute(int32(id)); err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Registro eliminado"})
}
