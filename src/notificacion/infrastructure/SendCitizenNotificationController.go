package infrastructure

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/notificacion/application"
)

type SendCitizenNotificationController struct {
	useCase *application.SendCitizenNotificationUseCase
}

func NewSendCitizenNotificationController(uc *application.SendCitizenNotificationUseCase) *SendCitizenNotificationController {
	return &SendCitizenNotificationController{useCase: uc}
}

func (c *SendCitizenNotificationController) Run(ctx *gin.Context) {
	var input application.SendCitizenNotificationInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	output := c.useCase.Execute(ctx.Request.Context(), &input)
	if output.Error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": output.Error.Error()})
		return
	}

	ctx.JSON(http.StatusOK, output)
}
