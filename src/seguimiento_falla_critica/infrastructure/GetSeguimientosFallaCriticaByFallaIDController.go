// GetSeguimientosFallaCriticaByFallaIDController.go
package infrastructure

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/seguimiento_falla_critica/application"
)

type GetSeguimientosFallaCriticaByFallaIDController struct {
	getByFallaIDUseCase *application.GetSeguimientosFallaCriticaByFallaIDUseCase
}

func NewGetSeguimientosFallaCriticaByFallaIDController(getByFallaIDUseCase *application.GetSeguimientosFallaCriticaByFallaIDUseCase) *GetSeguimientosFallaCriticaByFallaIDController {
	return &GetSeguimientosFallaCriticaByFallaIDController{
		getByFallaIDUseCase: getByFallaIDUseCase,
	}
}

func (ctrl *GetSeguimientosFallaCriticaByFallaIDController) Run(c *gin.Context) {
	fallaIDParam := c.Param("fallaId")
	fallaID, err := strconv.Atoi(fallaIDParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "ID de falla inválido",
			"error":   err.Error(),
		})
		return
	}

	seguimientos, err := ctrl.getByFallaIDUseCase.Run(int32(fallaID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "No se pudieron obtener los seguimientos para la falla",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, seguimientos)
}