package infrastructure

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/notificacion/application"
	"github.com/vicpoo/API_recolecta/src/notificacion/domain"
)

type NotificationRulesController struct {
	uc *application.ManageNotificationRulesUseCase
}

func NewNotificationRulesController(uc *application.ManageNotificationRulesUseCase) *NotificationRulesController {
	return &NotificationRulesController{uc: uc}
}

func (ctrl *NotificationRulesController) Upsert(c *gin.Context) {
	stateCode := c.Param("state_code")
	var req struct {
		Action        string `json:"action"`
		RadiusMeters  int    `json:"radius_meters"`
		Priority      int    `json:"priority"`
		Enabled       bool   `json:"enabled"`
		TemplateTitle string `json:"template_title"`
		TemplateBody  string `json:"template_body"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "payload inválido", "details": err.Error()})
		return
	}

	rule := &domain.NotificationRule{
		StateCode:     stateCode,
		Action:        req.Action,
		RadiusMeters:  req.RadiusMeters,
		Priority:      req.Priority,
		Enabled:       req.Enabled,
		TemplateTitle: req.TemplateTitle,
		TemplateBody:  req.TemplateBody,
	}

	if err := ctrl.uc.Upsert(c.Request.Context(), rule); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "regla guardada", "rule": rule})
}

func (ctrl *NotificationRulesController) GetByStateCode(c *gin.Context) {
	rule, err := ctrl.uc.GetByStateCode(c.Request.Context(), c.Param("state_code"))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"rule": rule})
}

func (ctrl *NotificationRulesController) List(c *gin.Context) {
	rules, err := ctrl.uc.List(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"rules": rules})
}

func (ctrl *NotificationRulesController) Delete(c *gin.Context) {
	if err := ctrl.uc.Delete(c.Request.Context(), c.Param("state_code")); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "regla eliminada"})
}
