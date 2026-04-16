package infrastructure

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/notificacion/application"
)

type AdminRealtimeSessionController struct {
	uc *application.ManageAdminRealtimeSessionUseCase
}

func NewAdminRealtimeSessionController(uc *application.ManageAdminRealtimeSessionUseCase) *AdminRealtimeSessionController {
	return &AdminRealtimeSessionController{uc: uc}
}

func (ctrl *AdminRealtimeSessionController) IssueUpgradeToken(c *gin.Context) {
	adminID := int32(c.GetInt("user_id"))
	output, err := ctrl.uc.IssueUpgradeToken(c.Request.Context(), adminID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, output)
}

func (ctrl *AdminRealtimeSessionController) ConsumeUpgradeToken(c *gin.Context) {
	var req struct {
		WSUpgradeToken string `json:"ws_upgrade_token"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "payload invalido", "details": err.Error()})
		return
	}

	output, err := ctrl.uc.ConsumeUpgradeToken(c.Request.Context(), req.WSUpgradeToken)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, output)
}

func (ctrl *AdminRealtimeSessionController) Heartbeat(c *gin.Context) {
	if err := ctrl.uc.Heartbeat(c.Request.Context(), c.Param("session_id")); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "heartbeat registrado"})
}

func (ctrl *AdminRealtimeSessionController) Disconnect(c *gin.Context) {
	if err := ctrl.uc.Disconnect(c.Request.Context(), c.Param("session_id")); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "session cerrada"})
}
