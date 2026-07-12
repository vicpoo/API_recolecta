package application

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/vicpoo/API_recolecta/src/core"
	"github.com/vicpoo/API_recolecta/src/notificacion/domain"
)

type ProcessTruckStateEventUseCase struct {
	rulesRepo domain.INotificationRuleRepository
	traceRepo domain.IEventTraceRepository
	sender    domain.PushNotificationSender
}

func NewProcessTruckStateEventUseCase(rulesRepo domain.INotificationRuleRepository, traceRepo domain.IEventTraceRepository, sender domain.PushNotificationSender) *ProcessTruckStateEventUseCase {
	return &ProcessTruckStateEventUseCase{rulesRepo: rulesRepo, traceRepo: traceRepo, sender: sender}
}

func (uc *ProcessTruckStateEventUseCase) Execute(ctx context.Context, event *domain.TruckStateEvent) error {
	hash := computeEventHash(event)

	acquired, err := uc.traceRepo.TryAcquireDeduplication(ctx, hash, event)
	if err != nil {
		return err
	}
	if !acquired {
		return nil
	}

	action := domain.ActionNotifyAdminOnly
	adminNotified := true
	citizenFanoutCount := 0

	rule, err := uc.rulesRepo.GetByStateCode(ctx, string(event.StateCode))
	if err == nil && rule.Enabled {
		// Si la regla está habilitada y el radio es > 0, buscamos domicilios cercanos
		if rule.Threshold > 0 && event.Lat != 0 && event.Lon != 0 {
			rdb, err := core.ConnectRedis()
			if err == nil {
				// Buscar domicilios en Redis domicilios:geo a rule.Threshold metros del camión
				locations, err := rdb.GeoRadius(ctx, "domicilios:geo", event.Lon, event.Lat, &redis.GeoRadiusQuery{
					Radius: float64(rule.Threshold),
					Unit:   "m",
				}).Result()

				if err == nil && len(locations) > 0 {
					todayStr := time.Now().Format("2006-01-02")
					userTokens := make(map[string]string)

					for _, loc := range locations {
						domicilioID := loc.Name

						// Obtener el ciudadano_id asociado a este domicilio en Redis
						domKey := fmt.Sprintf("domicilio:%s", domicilioID)
						citizenID, err := rdb.HGet(ctx, domKey, "ciudadano_id").Result()
						if err != nil || citizenID == "" {
							continue
						}

						// Control de idempotencia diario: evitar spam por ciudadano
						dupKey := fmt.Sprintf("notification:sent:%s:%d:%s", citizenID, event.TruckID, todayStr)
						alreadySent, err := rdb.SIsMember(ctx, dupKey, string(event.StateCode)).Result()
						if err != nil || alreadySent {
							continue
						}

						// Obtener el token FCM del ciudadano
						userKey := fmt.Sprintf("user:%s", citizenID)
						token, err := rdb.HGet(ctx, userKey, "fcm_token").Result()
						if err == nil && token != "" {
							userTokens[citizenID] = token
							// Registrar envío para control de duplicidad
							rdb.SAdd(ctx, dupKey, string(event.StateCode))
							rdb.Expire(ctx, dupKey, 24*time.Hour)
						}
					}

					// Si hay tokens a notificar, enviamos la notificación push
					if len(userTokens) > 0 {
						citizenFanoutCount = len(userTokens)
						action = domain.Action("NOTIFY_CITIZENS_" + event.StateCode)
						
						notification := &domain.PushNotification{
							Title: fmt.Sprintf("Camión de Basura cerca de ti"),
							Body:  fmt.Sprintf("El camión %d cambió a estado: %s", event.TruckID, event.StateCode),
							Type:  string(event.StateCode),
							Data: map[string]string{
								"truck_id":   fmt.Sprintf("%d", event.TruckID),
								"state_code": string(event.StateCode),
							},
						}

						// Enviar notificaciones
						_, _ = uc.sender.Send(ctx, userTokens, notification)
					}
				}
			}
		}
	} else if err == nil && !rule.Enabled {
		adminNotified = false
	}

	trace := &domain.EventTraceRecord{
		EventID:            event.EventID,
		EventHash:          hash,
		EventType:          event.EventType,
		EventVersion:       event.EventVersion,
		TruckID:            event.TruckID,
		StateCode:          event.StateCode,
		ResolvedAction:     action,
		AdminNotified:      adminNotified,
		CitizenFanoutCount: citizenFanoutCount,
		Result:             "processed",
		CreatedAt:          event.OccurredAt,
	}
	return uc.traceRepo.SaveTrace(ctx, trace)
}

func computeEventHash(event *domain.TruckStateEvent) string {
	raw := fmt.Sprintf("%s:%d:%s:%s", event.EventID, event.TruckID, event.StateCode, event.EventType)
	h := sha256.Sum256([]byte(raw))
	return hex.EncodeToString(h[:])
}
