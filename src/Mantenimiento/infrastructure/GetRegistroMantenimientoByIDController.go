//GetRegistroMantenimientoByIDController.go
package infrastructure

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/Mantenimiento/application"
)

type GetRegistroMantenimientoByIDController struct {
	getByIDUseCase *application.GetRegistroMantenimientoByIDUseCase
}

func NewGetRegistroMantenimientoByIDController(getByIDUseCase *application.GetRegistroMantenimientoByIDUseCase) *GetRegistroMantenimientoByIDController {
	return &GetRegistroMantenimientoByIDController{
		getByIDUseCase: getByIDUseCase,
	}
}

// @Summary      Registro por ID
// @Tags         RegistroMantenimiento
// @Produce      json
// @Param        id path int true "ID"
// @Success      200 {object} map[string]interface{}
// @Failure      400 {object} map[string]interface{}
// @Router       /api/registros-mantenimiento/{id} [get]
func (ctrl *GetRegistroMantenimientoByIDController) Run(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "ID inválido",
			"error":   err.Error(),
		})
		return
	}

	registro, err := ctrl.getByIDUseCase.Run(int32(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "No se pudo encontrar el registro de mantenimiento",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, registro)
}