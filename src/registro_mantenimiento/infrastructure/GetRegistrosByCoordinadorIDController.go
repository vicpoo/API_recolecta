// GetRegistrosByCoordinadorIDController.go
package infrastructure

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/registro_mantenimiento/application"
)

type GetRegistrosByCoordinadorIDController struct {
	getByCoordinadorIDUseCase *application.GetRegistrosByCoordinadorIDUseCase
}

func NewGetRegistrosByCoordinadorIDController(getByCoordinadorIDUseCase *application.GetRegistrosByCoordinadorIDUseCase) *GetRegistrosByCoordinadorIDController {
	return &GetRegistrosByCoordinadorIDController{
		getByCoordinadorIDUseCase: getByCoordinadorIDUseCase,
	}
}

func (ctrl *GetRegistrosByCoordinadorIDController) Run(c *gin.Context) {
	coordinadorIDParam := c.Param("coordinador_id")
	coordinadorID, err := strconv.Atoi(coordinadorIDParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "ID de coordinador inválido",
			"error":   err.Error(),
		})
		return
	}

	registros, err := ctrl.getByCoordinadorIDUseCase.Run(int32(coordinadorID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "No se pudieron obtener los registros de mantenimiento del coordinador",
			"error":   err.Error(),
		})
		return
	}

	if len(registros) == 0 {
		c.JSON(http.StatusOK, gin.H{
			"message": "No se encontraron registros de mantenimiento para este coordinador",
			"data":    []string{},
		})
		return
	}

	c.JSON(http.StatusOK, registros)
}