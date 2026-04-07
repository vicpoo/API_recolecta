package infrastructure

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/notificacion/application"
	"github.com/vicpoo/API_recolecta/src/notificacion/domain"
)

type SendCitizenNotificationController struct {
	useCase *application.SendCitizenNotificationUseCase
}

type SendCitizenNotificationResponse struct {
	Success bool                       `json:"success"`
	Message string                     `json:"message"`
	Results map[string]domain.SendResult `json:"results"`
	Error   string                     `json:"error,omitempty"`
}

func NewSendCitizenNotificationController(uc *application.SendCitizenNotificationUseCase) *SendCitizenNotificationController {
	return &SendCitizenNotificationController{useCase: uc}
}

func (c *SendCitizenNotificationController) Run(ctx *gin.Context) {
	var input application.SendCitizenNotificationInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, SendCitizenNotificationResponse{
			Success: false,
			Message: "Solicitud invalida",
			Results: map[string]domain.SendResult{},
			Error:   err.Error(),
		})
		return
	}

	output := c.useCase.Execute(ctx.Request.Context(), &input)
	if output.Error != nil {
		ctx.JSON(http.StatusBadRequest, SendCitizenNotificationResponse{
			Success: false,
			Message: "No se pudo procesar el envio",
			Results: output.Results,
			Error:   output.Error.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, SendCitizenNotificationResponse{
		Success: true,
		Message: "Notificaciones procesadas",
		Results: output.Results,
	})
}
