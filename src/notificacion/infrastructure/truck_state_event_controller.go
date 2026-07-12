package infrastructure

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/notificacion/application"
	"github.com/vicpoo/API_recolecta/src/notificacion/domain"
)

type TruckStateEventController struct {
	processUc *application.ProcessTruckStateEventUseCase
	queryUc   *application.QueryEventTraceUseCase
}

func NewTruckStateEventController(processUc *application.ProcessTruckStateEventUseCase, queryUc *application.QueryEventTraceUseCase) *TruckStateEventController {
	return &TruckStateEventController{processUc: processUc, queryUc: queryUc}
}

type processTruckStateRequest struct {
	EventID      string   `json:"event_id"`
	EventType    string   `json:"event_type"`
	EventVersion string   `json:"event_version"`
	TruckID      int32    `json:"truck_id"`
	StateCode    string   `json:"state_code"`
	Lat          float64  `json:"lat"`
	Lon          float64  `json:"lon"`
	Payload      *struct {
		StateCode string  `json:"state_code"`
		Lat       float64 `json:"lat"`
		Lon       float64 `json:"lon"`
	} `json:"payload"`
}

func (ctrl *TruckStateEventController) Process(c *gin.Context) {
	var req processTruckStateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	lat := req.Lat
	lon := req.Lon
	stateCode := req.StateCode
	if req.Payload != nil {
		lat = req.Payload.Lat
		lon = req.Payload.Lon
		if req.Payload.StateCode != "" {
			stateCode = req.Payload.StateCode
		}
	}

	event := &domain.TruckStateEvent{
		EventID:      req.EventID,
		EventType:    req.EventType,
		EventVersion: domain.EventVersion(req.EventVersion),
		TruckID:      req.TruckID,
		StateCode:    domain.StateCode(stateCode),
		OccurredAt:   time.Now().UTC(),
		Lat:          lat,
		Lon:          lon,
	}
	if err := ctrl.processUc.Execute(c.Request.Context(), event); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "processed"})
}

func (ctrl *TruckStateEventController) GetByEventID(c *gin.Context) {
	eventID := c.Param("event_id")
	trace, err := ctrl.queryUc.GetByEventID(c.Request.Context(), eventID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, trace)
}

func (ctrl *TruckStateEventController) ListByTruckID(c *gin.Context) {
	truckIDStr := c.Param("truck_id")
	truckID, err := strconv.ParseInt(truckIDStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid truck_id"})
		return
	}
	traces, err := ctrl.queryUc.ListByTruckID(c.Request.Context(), int32(truckID), 100)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, traces)
}
