package infrastructure

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/vicpoo/API_recolecta/src/notificacion/application"
)

type AdminRealtimeSessionController struct {
	uc *application.ManageAdminRealtimeSessionUseCase
}

func NewAdminRealtimeSessionController(uc *application.ManageAdminRealtimeSessionUseCase) *AdminRealtimeSessionController {
	return &AdminRealtimeSessionController{uc: uc}
}

func (ctrl *AdminRealtimeSessionController) IssueUpgradeToken(c *gin.Context) {
	adminIDRaw, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "user_id not found in context"})
		return
	}
	adminID, ok := adminIDRaw.(int32)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "invalid user_id type"})
		return
	}
	sessionID := uuid.New().String()
	claim, err := ctrl.uc.IssueUpgradeToken(c.Request.Context(), adminID, sessionID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"jti":        claim.JTI,
		"session_id": claim.SessionID,
		"expires_at": claim.ExpiresAt,
	})
}

func (ctrl *AdminRealtimeSessionController) ConsumeUpgradeToken(c *gin.Context) {
	var body struct {
		JTI string `json:"jti"`
	}
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	claim, err := ctrl.uc.ConsumeUpgradeToken(c.Request.Context(), body.JTI)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	session, err := ctrl.uc.RegisterSession(c.Request.Context(), claim)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, session)
}

func (ctrl *AdminRealtimeSessionController) Heartbeat(c *gin.Context) {
	sessionID := c.Param("session_id")
	if err := ctrl.uc.Heartbeat(c.Request.Context(), sessionID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}

func (ctrl *AdminRealtimeSessionController) GetSession(c *gin.Context) {
	sessionID := c.Param("session_id")
	session, err := ctrl.uc.GetSession(c.Request.Context(), sessionID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, session)
}

func (ctrl *AdminRealtimeSessionController) Disconnect(c *gin.Context) {
	sessionID := c.Param("session_id")
	if err := ctrl.uc.Disconnect(c.Request.Context(), sessionID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}
