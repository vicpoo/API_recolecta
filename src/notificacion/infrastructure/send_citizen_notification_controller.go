package infrastructure

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/notificacion/application"
	"github.com/vicpoo/API_recolecta/src/notificacion/domain"
)

type SendCitizenNotificationController struct {
	uc *application.SendCitizenNotificationUseCase
}

func NewSendCitizenNotificationController(uc *application.SendCitizenNotificationUseCase) *SendCitizenNotificationController {
	return &SendCitizenNotificationController{uc: uc}
}

type sendCitizenRequest struct {
	UserIDs []string          `json:"user_ids"`
	Title   string            `json:"title"`
	Body    string            `json:"body"`
	Type    string            `json:"type"`
	Data    map[string]string `json:"data"`
}

func (ctrl *SendCitizenNotificationController) Run(c *gin.Context) {
	var req sendCitizenRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	notification := &domain.PushNotification{
		Title: req.Title,
		Body:  req.Body,
		Type:  req.Type,
		Data:  req.Data,
	}

	results, err := ctrl.uc.Execute(c.Request.Context(), req.UserIDs, notification)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, results)
}
