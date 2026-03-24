package ciudadano_redis_test

import (
	"context"
	"testing"

	"github.com/alicebob/miniredis/v2"
	goredis "github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	ciudadanoredis "github.com/vicpoo/API_recolecta/src/ciudadano/infrastructure/redis"
	"github.com/vicpoo/API_recolecta/src/ciudadano/domain/ports"
)

// setupRepo spins up a fresh miniredis instance and returns the repository and
// the underlying miniredis handle for raw inspection in assertions.
// The miniredis server and Redis client are automatically closed via t.Cleanup.
func setupRepo(t *testing.T) (*miniredis.Miniredis, ports.CiudadanoRepository) {
	t.Helper()
	mr := miniredis.RunT(t)
	rdb := goredis.NewClient(&goredis.Options{Addr: mr.Addr()})
	t.Cleanup(func() { _ = rdb.Close() })
	repo := ciudadanoredis.NewRedisCiudadanoRepository(rdb)
	return mr, repo
}

// ─── RegisterUser ─────────────────────────────────────────────────────────────

func TestRegisterUser_StoresGeoAndHash(t *testing.T) {
	mr, repo := setupRepo(t)
	ctx := context.Background()

	err := repo.RegisterUser(ctx, "u1", 13.361389, 38.115556, "token-abc")
	require.NoError(t, err)

	// FCM token persisted in HASH
	assert.Equal(t, "token-abc", mr.HGet("user:u1", "fcm_token"))

	// updated_at is set
	assert.NotEmpty(t, mr.HGet("user:u1", "updated_at"))

	// GEO position persisted
	lon, lat, err := repo.GetUserGeo(ctx, "u1")
	require.NoError(t, err)
	assert.InDelta(t, 13.361389, lon, 0.001)
	assert.InDelta(t, 38.115556, lat, 0.001)
}

// ─── UpdateUserFCMToken ───────────────────────────────────────────────────────

func TestUpdateUserFCMToken_OnlyTokenChanges(t *testing.T) {
	mr, repo := setupRepo(t)
	ctx := context.Background()

	require.NoError(t, repo.RegisterUser(ctx, "u1", 13.361389, 38.115556, "old-token"))

	require.NoError(t, repo.UpdateUserFCMToken(ctx, "u1", "new-token"))

	assert.Equal(t, "new-token", mr.HGet("user:u1", "fcm_token"))

	// Coordinates should be unchanged
	lon, lat, err := repo.GetUserGeo(ctx, "u1")
	require.NoError(t, err)
	assert.InDelta(t, 13.361389, lon, 0.001)
	assert.InDelta(t, 38.115556, lat, 0.001)
}

// ─── UpdateUserGeoCoordinates ─────────────────────────────────────────────────

func TestUpdateUserGeoCoordinates_OnlyCoordsChange(t *testing.T) {
	mr, repo := setupRepo(t)
	ctx := context.Background()

	require.NoError(t, repo.RegisterUser(ctx, "u1", 13.361389, 38.115556, "token"))

	require.NoError(t, repo.UpdateUserGeoCoordinates(ctx, "u1", 15.087269, 37.502669))

	lon, lat, err := repo.GetUserGeo(ctx, "u1")
	require.NoError(t, err)
	assert.InDelta(t, 15.087269, lon, 0.001)
	assert.InDelta(t, 37.502669, lat, 0.001)

	// FCM token should be unchanged
	assert.Equal(t, "token", mr.HGet("user:u1", "fcm_token"))
}

// ─── GetUserFCMToken ──────────────────────────────────────────────────────────

func TestGetUserFCMToken_ReturnsStoredToken(t *testing.T) {
	_, repo := setupRepo(t)
	ctx := context.Background()

	require.NoError(t, repo.RegisterUser(ctx, "u1", 13.361389, 38.115556, "my-token"))

	token, err := repo.GetUserFCMToken(ctx, "u1")
	require.NoError(t, err)
	assert.Equal(t, "my-token", token)
}

func TestGetUserFCMToken_ReturnsEmptyForMissingUser(t *testing.T) {
	_, repo := setupRepo(t)
	ctx := context.Background()

	token, err := repo.GetUserFCMToken(ctx, "nonexistent")
	require.NoError(t, err)
	assert.Equal(t, "", token)
}

// ─── GetAllUserIDs ────────────────────────────────────────────────────────────

func TestGetAllUserIDs_ReturnsAllRegisteredUsers(t *testing.T) {
	_, repo := setupRepo(t)
	ctx := context.Background()

	require.NoError(t, repo.RegisterUser(ctx, "u1", 13.361389, 38.115556, "t1"))
	require.NoError(t, repo.RegisterUser(ctx, "u2", 15.087269, 37.502669, "t2"))

	ids, err := repo.GetAllUserIDs(ctx)
	require.NoError(t, err)
	assert.ElementsMatch(t, []string{"u1", "u2"}, ids)
}

// ─── HasFCMToken ──────────────────────────────────────────────────────────────

func TestHasFCMToken_TrueWhenTokenExists(t *testing.T) {
	_, repo := setupRepo(t)
	ctx := context.Background()

	require.NoError(t, repo.RegisterUser(ctx, "u1", 13.361389, 38.115556, "token"))

	has, err := repo.HasFCMToken(ctx, "u1")
	require.NoError(t, err)
	assert.True(t, has)
}

func TestHasFCMToken_FalseWhenUserAbsent(t *testing.T) {
	_, repo := setupRepo(t)
	ctx := context.Background()

	has, err := repo.HasFCMToken(ctx, "nonexistent")
	require.NoError(t, err)
	assert.False(t, has)
}

// ─── GetUsersInRadius ─────────────────────────────────────────────────────────

func TestGetUsersInRadius_FiltersCorrectly(t *testing.T) {
	_, repo := setupRepo(t)
	ctx := context.Background()

	// Palermo (Sicily): well-known Redis example coordinates
	require.NoError(t, repo.RegisterUser(ctx, "u1", 13.361389, 38.115556, "t1"))
	// Catania (Sicily): ~166 km from Palermo
	require.NoError(t, repo.RegisterUser(ctx, "u2", 15.087269, 37.502669, "t2"))

	// 100 km radius from Palermo — only u1 should match
	near, err := repo.GetUsersInRadius(ctx, 13.361389, 38.115556, 100_000)
	require.NoError(t, err)
	assert.Contains(t, near, "u1")
	assert.NotContains(t, near, "u2")

	// 200 km radius — both should match
	all, err := repo.GetUsersInRadius(ctx, 13.361389, 38.115556, 200_000)
	require.NoError(t, err)
	assert.Contains(t, all, "u1")
	assert.Contains(t, all, "u2")
}
