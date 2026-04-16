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

func (ctrl *NotificationRulesController) List(c *gin.Context) {
	rules, err := ctrl.uc.List(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, rules)
}

func (ctrl *NotificationRulesController) GetByStateCode(c *gin.Context) {
	code := c.Param("state_code")
	rule, err := ctrl.uc.GetByStateCode(c.Request.Context(), code)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, rule)
}

func (ctrl *NotificationRulesController) Upsert(c *gin.Context) {
	code := c.Param("state_code")
	var rule domain.NotificationRule
	if err := c.ShouldBindJSON(&rule); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	rule.StateCode = code
	if err := ctrl.uc.Upsert(c.Request.Context(), rule); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, rule)
}

func (ctrl *NotificationRulesController) Delete(c *gin.Context) {
	code := c.Param("state_code")
	if err := ctrl.uc.Delete(c.Request.Context(), code); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}
