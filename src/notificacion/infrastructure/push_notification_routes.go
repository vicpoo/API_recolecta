package infrastructure

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/config"
	"github.com/vicpoo/API_recolecta/src/core"
	"github.com/vicpoo/API_recolecta/src/notificacion/application"
)

type PushNotificationRouter struct {
	engine *gin.Engine
}

func NewPushNotificationRouter(engine *gin.Engine) *PushNotificationRouter {
	return &PushNotificationRouter{engine: engine}
}

func (r *PushNotificationRouter) Run() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("no se pudo cargar la configuración: %v", err)
	}

	fcmClient, err := NewFCMClient(cfg.FCMCredentialsFile)
	if err != nil {
		log.Fatalf("no se pudo inicializar el cliente FCM: %v", err)
	}

	redisRepo := NewRedisNotificationRepository(core.GetRedis())
	uc := application.NewSendCitizenNotificationUseCase(fcmClient, redisRepo)
	ctrl := NewSendCitizenNotificationController(uc)
	rulesRepo := NewRedisNotificationRuleRepository(core.GetRedis())
	rulesUc := application.NewManageNotificationRulesUseCase(rulesRepo)
	rulesCtrl := NewNotificationRulesController(rulesUc)
	traceRepo := NewRedisEventTraceRepository(core.GetRedis())
	processEventUc := application.NewProcessTruckStateEventUseCase(rulesRepo, traceRepo)
	queryEventTraceUc := application.NewQueryEventTraceUseCase(traceRepo)
	processEventCtrl := NewTruckStateEventController(processEventUc, queryEventTraceUc)
	adminRealtimeRepo := NewRedisAdminRealtimeSessionRepository(core.GetRedis())
	adminRealtimeUc := application.NewManageAdminRealtimeSessionUseCase(adminRealtimeRepo)
	adminRealtimeCtrl := NewAdminRealtimeSessionController(adminRealtimeUc)
	observabilityUc := application.NewGetObservabilitySummaryUseCase(traceRepo, adminRealtimeRepo)
	observabilityCtrl := NewObservabilityController(observabilityUc)

	group := r.engine.Group("/api/notifications")
	{
		group.POST("/citizens/send", ctrl.Run)
		group.POST("/events/truck-state", processEventCtrl.Process)
		group.GET("/events/traces/:event_id", processEventCtrl.GetByEventID)
		group.GET("/events/traces/truck/:truck_id", processEventCtrl.ListByTruckID)
	}

	rulesGroup := r.engine.Group("/api/notifications/rules", core.JWTAuthMiddleware())
	{
		rulesGroup.GET("", rulesCtrl.List)
		rulesGroup.GET("/:state_code", rulesCtrl.GetByStateCode)
		rulesGroup.PUT("/:state_code", rulesCtrl.Upsert)
		rulesGroup.DELETE("/:state_code", rulesCtrl.Delete)
	}

	observabilityGroup := r.engine.Group("/api/notifications/observability", core.JWTAuthMiddleware(), core.RequireRole(core.ADMIN))
	{
		observabilityGroup.GET("/:truck_id", observabilityCtrl.Summary)
	}

	realtimeGroup := r.engine.Group("/api/realtime/ws")
	{
		realtimeGroup.POST("/upgrade-token", core.JWTAuthMiddleware(), core.RequireRole(core.ADMIN), adminRealtimeCtrl.IssueUpgradeToken)
		realtimeGroup.POST("/sessions/consume", adminRealtimeCtrl.ConsumeUpgradeToken)
		realtimeGroup.POST("/sessions/:session_id/heartbeat", core.JWTAuthMiddleware(), core.RequireRole(core.ADMIN), adminRealtimeCtrl.Heartbeat)
		realtimeGroup.GET("/sessions/:session_id", core.JWTAuthMiddleware(), core.RequireRole(core.ADMIN), adminRealtimeCtrl.GetSession)
		realtimeGroup.DELETE("/sessions/:session_id", core.JWTAuthMiddleware(), core.RequireRole(core.ADMIN), adminRealtimeCtrl.Disconnect)
	}
}
