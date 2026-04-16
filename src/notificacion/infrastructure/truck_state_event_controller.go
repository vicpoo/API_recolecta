package infrastructure

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/notificacion/application"
	"github.com/vicpoo/API_recolecta/src/notificacion/domain"
)

type TruckStateEventController struct {
	processUc *application.ProcessTruckStateEventUseCase
	queryUc   *application.QueryEventTraceUseCase
}

func NewTruckStateEventController(
	processUc *application.ProcessTruckStateEventUseCase,
	queryUc *application.QueryEventTraceUseCase,
) *TruckStateEventController {
	return &TruckStateEventController{
		processUc: processUc,
		queryUc:   queryUc,
	}
}

func (ctrl *TruckStateEventController) Process(c *gin.Context) {
	var event domain.TruckStateEvent
	if err := c.ShouldBindJSON(&event); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "payload invalido", "details": err.Error()})
		return
	}

	output, err := ctrl.processUc.Execute(c.Request.Context(), &event)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, output)
}

func (ctrl *TruckStateEventController) GetByEventID(c *gin.Context) {
	trace, err := ctrl.queryUc.GetByEventID(c.Request.Context(), c.Param("event_id"))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"trace": trace})
}

func (ctrl *TruckStateEventController) ListByTruckID(c *gin.Context) {
	truckIDRaw := c.Param("truck_id")
	truckIDParsed, err := strconv.ParseInt(truckIDRaw, 10, 32)
	if err != nil || truckIDParsed <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "truck_id invalido"})
		return
	}

	limit := int64(20)
	if limitRaw := c.Query("limit"); limitRaw != "" {
		parsedLimit, parseErr := strconv.ParseInt(limitRaw, 10, 64)
		if parseErr != nil || parsedLimit <= 0 || parsedLimit > 100 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "limit invalido, rango permitido 1..100"})
			return
		}
		limit = parsedLimit
	}

	traces, err := ctrl.queryUc.ListByTruckID(c.Request.Context(), int32(truckIDParsed), limit)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"traces": traces})
}
