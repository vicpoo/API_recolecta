package domain

import "time"

const (
	EventVersionV1 = "v1"
)

const (
	StateInit          = "INIT"
	StateInRoute       = "IN_ROUTE"
	StateWarn          = "WARN"
	StateArrival       = "ARRIVAL"
	StateDeparture     = "DEPARTURE"
	StateComeback      = "COMEBACK"
	StateEmergency     = "EMERGENCY"
	StateInoperability = "INOPERABILITY"
)

const (
	ActionNotifyAdminOnly          = "NOTIFY_ADMIN_ONLY"
	ActionNotifyAdminAndCitizens   = "NOTIFY_ADMIN_AND_CITIZENS"
)

// TruckStateEvent is the server-owned input contract for mobility clients.
// Mobile clients must respect this payload and version.
type TruckStateEvent struct {
	EventID      string                 `json:"event_id"`
	EventType    string                 `json:"event_type"`
	EventVersion string                 `json:"event_version"`
	TruckID      int32                  `json:"truck_id"`
	OccurredAt   time.Time              `json:"occurred_at"`
	Payload      map[string]interface{} `json:"payload"`
}

// AdminWSUpgradeTokenClaim represents a one-purpose token used only for
// upgrading an authenticated admin session into a realtime websocket channel.
type AdminWSUpgradeTokenClaim struct {
	JTI         string    `json:"jti"`
	AdminID     int32     `json:"admin_id"`
	SessionID   string    `json:"session_id"`
	ServerEpoch string    `json:"server_epoch"`
	IssuedAt    time.Time `json:"issued_at"`
	ExpiresAt   time.Time `json:"expires_at"`
}
