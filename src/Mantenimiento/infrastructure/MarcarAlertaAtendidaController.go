//MarcarAlertaAtendidaController.go
package infrastructure

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/Mantenimiento/application"
)

type MarcarAlertaAtendidaController struct {
	marcarAtendidaUseCase *application.MarcarAlertaAtendidaUseCase
}

func NewMarcarAlertaAtendidaController(marcarAtendidaUseCase *application.MarcarAlertaAtendidaUseCase) *MarcarAlertaAtendidaController {
	return &MarcarAlertaAtendidaController{
		marcarAtendidaUseCase: marcarAtendidaUseCase,
	}
}

// @Summary      Marcar alerta como atendida
// @Tags         AlertaMantenimiento
// @Produce      json
// @Param        id path int true "ID"
// @Success      200 {object} map[string]interface{}
// @Failure      400 {object} map[string]interface{}
// @Router       /api/alertas-mantenimiento/{id}/atender [patch]
func (ctrl *MarcarAlertaAtendidaController) Run(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "ID inválido",
			"error":   err.Error(),
		})
		return
	}

	err = ctrl.marcarAtendidaUseCase.Run(int32(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "No se pudo marcar la alerta como atendida",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Alerta marcada como atendida exitosamente",
	})
}