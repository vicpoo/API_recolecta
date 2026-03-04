package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/Rutas/application"
)

type ExistsRegistroVaciadoController struct {
	uc *application.ExistsRegistroVaciadoUseCase
}

func NewExistsRegistroVaciadoController(
	uc *application.ExistsRegistroVaciadoUseCase,
) *ExistsRegistroVaciadoController {
	return &ExistsRegistroVaciadoController{uc: uc}
}

// @Summary      Verificar existencia registro vaciado
// @Tags         RegistroVaciado
// @Produce      json
// @Param        id path int true "ID"
// @Success      200 {object} map[string]interface{}
// @Router       /api/registro-vaciado/exists/{id} [get]
func (c *ExistsRegistroVaciadoController) Run(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "id inválido"})
		return
	}

	exists, err := c.uc.Execute(int32(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"exists": exists,
	})
}
