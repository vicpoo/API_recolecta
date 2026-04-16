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

	group := r.engine.Group("/api/notifications")
	{
		group.POST("/citizens/send", ctrl.Run)
	}
}
