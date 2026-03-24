package ports

import "context"

type CiudadanoRepository interface {
	// RegisterUser registers or overwrites a citizen's geo position and FCM token.
	RegisterUser(ctx context.Context, userID string, longitude, latitude float64, fcmToken string) error

	// UpdateUserFCMToken updates only the FCM token for an existing user.
	UpdateUserFCMToken(ctx context.Context, userID string, fcmToken string) error

	// UpdateUserGeoCoordinates updates only the geospatial coordinates for an existing user.
	UpdateUserGeoCoordinates(ctx context.Context, userID string, longitude, latitude float64) error

	// GetUserGeo returns the stored longitude and latitude for a user.
	GetUserGeo(ctx context.Context, userID string) (longitude, latitude float64, err error)

	// GetUserFCMToken returns the FCM token stored for a user. Returns ("", nil) if not found.
	GetUserFCMToken(ctx context.Context, userID string) (string, error)

	// GetAllUserIDs returns every user_id present in the geospatial index.
	GetAllUserIDs(ctx context.Context) ([]string, error)

	// HasFCMToken reports whether the user has a non-empty FCM token stored.
	HasFCMToken(ctx context.Context, userID string) (bool, error)

	// GetUsersInRadius returns user IDs within radiusMeters of the given coordinates.
	GetUsersInRadius(ctx context.Context, longitude, latitude float64, radiusMeters float64) ([]string, error)
}
