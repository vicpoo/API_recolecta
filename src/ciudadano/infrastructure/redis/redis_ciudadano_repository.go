package redis

import (
	"context"
	"fmt"
	"time"

	goredis "github.com/redis/go-redis/v9"
	"github.com/vicpoo/API_recolecta/src/ciudadano/domain/ports"
)

const (
	usersGeoKey    = "users:geo"
	userHashKeyFmt = "user:%s"
)

type RedisCiudadanoRepository struct {
	rdb *goredis.Client
}

// NewRedisCiudadanoRepository creates a new adapter. The caller provides the
// Redis client so the repository is easily testable with miniredis.
func NewRedisCiudadanoRepository(rdb *goredis.Client) ports.CiudadanoRepository {
	return &RedisCiudadanoRepository{rdb: rdb}
}

func (r *RedisCiudadanoRepository) RegisterUser(ctx context.Context, userID string, longitude, latitude float64, fcmToken string) error {
	pipe := r.rdb.TxPipeline()
	pipe.GeoAdd(ctx, usersGeoKey, &goredis.GeoLocation{
		Name:      userID,
		Longitude: longitude,
		Latitude:  latitude,
	})
	pipe.HSet(ctx, fmt.Sprintf(userHashKeyFmt, userID),
		"fcm_token", fcmToken,
		"updated_at", time.Now().UTC().Format(time.RFC3339),
	)
	_, err := pipe.Exec(ctx)
	return err
}

func (r *RedisCiudadanoRepository) UpdateUserFCMToken(ctx context.Context, userID string, fcmToken string) error {
	return r.rdb.HSet(ctx, fmt.Sprintf(userHashKeyFmt, userID),
		"fcm_token", fcmToken,
		"updated_at", time.Now().UTC().Format(time.RFC3339),
	).Err()
}

func (r *RedisCiudadanoRepository) UpdateUserGeoCoordinates(ctx context.Context, userID string, longitude, latitude float64) error {
	pipe := r.rdb.TxPipeline()
	pipe.GeoAdd(ctx, usersGeoKey, &goredis.GeoLocation{
		Name:      userID,
		Longitude: longitude,
		Latitude:  latitude,
	})
	pipe.HSet(ctx, fmt.Sprintf(userHashKeyFmt, userID),
		"updated_at", time.Now().UTC().Format(time.RFC3339),
	)
	_, err := pipe.Exec(ctx)
	return err
}

func (r *RedisCiudadanoRepository) GetUserGeo(ctx context.Context, userID string) (float64, float64, error) {
	positions, err := r.rdb.GeoPos(ctx, usersGeoKey, userID).Result()
	if err != nil {
		return 0, 0, err
	}
	if len(positions) == 0 || positions[0] == nil {
		return 0, 0, fmt.Errorf("user %q not found in geo index", userID)
	}
	return positions[0].Longitude, positions[0].Latitude, nil
}

func (r *RedisCiudadanoRepository) GetUserFCMToken(ctx context.Context, userID string) (string, error) {
	val, err := r.rdb.HGet(ctx, fmt.Sprintf(userHashKeyFmt, userID), "fcm_token").Result()
	if err == goredis.Nil {
		return "", nil
	}
	return val, err
}

func (r *RedisCiudadanoRepository) GetAllUserIDs(ctx context.Context) ([]string, error) {
	return r.rdb.ZRange(ctx, usersGeoKey, 0, -1).Result()
}

func (r *RedisCiudadanoRepository) HasFCMToken(ctx context.Context, userID string) (bool, error) {
	val, err := r.rdb.HGet(ctx, fmt.Sprintf(userHashKeyFmt, userID), "fcm_token").Result()
	if err == goredis.Nil {
		return false, nil
	}
	if err != nil {
		return false, err
	}
	return val != "", nil
}

func (r *RedisCiudadanoRepository) GetUsersInRadius(ctx context.Context, longitude, latitude float64, radiusMeters float64) ([]string, error) {
	locations, err := r.rdb.GeoRadius(ctx, usersGeoKey, longitude, latitude, &goredis.GeoRadiusQuery{
		Radius:      radiusMeters,
		Unit:        "m",
		Sort:        "ASC",
		Count:       0,
		Store:       "",
		StoreDist:   "",
		WithCoord:   false,
		WithDist:    false,
		WithGeoHash: false,
	}).Result()
	if err != nil {
		return nil, err
	}

	userIDs := make([]string, 0, len(locations))
	for _, location := range locations {
		userIDs = append(userIDs, location.Name)
	}

	return userIDs, nil
}
