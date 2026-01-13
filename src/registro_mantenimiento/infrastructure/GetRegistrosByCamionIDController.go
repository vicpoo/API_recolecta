//GetRegistrosByCamionIDController.go
package infrastructure

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/registro_mantenimiento/application"
)

type GetRegistrosByCamionIDController struct {
	getByCamionIDUseCase *application.GetRegistrosByCamionIDUseCase
}

func NewGetRegistrosByCamionIDController(getByCamionIDUseCase *application.GetRegistrosByCamionIDUseCase) *GetRegistrosByCamionIDController {
	return &GetRegistrosByCamionIDController{
		getByCamionIDUseCase: getByCamionIDUseCase,
	}
}

func (ctrl *GetRegistrosByCamionIDController) Run(c *gin.Context) {
	camionIDParam := c.Param("camion_id")
	camionID, err := strconv.Atoi(camionIDParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "ID de camión inválido",
			"error":   err.Error(),
		})
		return
	}

	registros, err := ctrl.getByCamionIDUseCase.Run(int32(camionID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "No se pudieron obtener los registros de mantenimiento del camión",
			"error":   err.Error(),
		})
		return
	}

	if len(registros) == 0 {
		c.JSON(http.StatusOK, gin.H{
			"message": "No se encontraron registros de mantenimiento para este camión",
			"data":    []string{},
		})
		return
	}

	c.JSON(http.StatusOK, registros)
}