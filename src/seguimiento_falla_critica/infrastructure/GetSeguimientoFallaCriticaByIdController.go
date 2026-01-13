// GetSeguimientoFallaCriticaByIdController.go
package infrastructure

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/seguimiento_falla_critica/application"
)

type GetSeguimientoFallaCriticaByIdController struct {
	getByIdUseCase *application.GetSeguimientoFallaCriticaByIdUseCase
}

func NewGetSeguimientoFallaCriticaByIdController(getByIdUseCase *application.GetSeguimientoFallaCriticaByIdUseCase) *GetSeguimientoFallaCriticaByIdController {
	return &GetSeguimientoFallaCriticaByIdController{
		getByIdUseCase: getByIdUseCase,
	}
}

func (ctrl *GetSeguimientoFallaCriticaByIdController) Run(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "ID inválido",
			"error":   err.Error(),
		})
		return
	}

	seguimiento, err := ctrl.getByIdUseCase.Run(int32(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "No se pudo obtener el seguimiento de falla crítica",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, seguimiento)
}