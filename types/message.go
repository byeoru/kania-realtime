package types

type Response struct {
	Title string `json:"title"`
	Body  any    `json:"body"`
}

type UpdateSector struct {
	Sector     int32  `json:"sector"`
	OldRealmID int64  `json:"old_realm_id"`
	NewRealmID int64  `json:"new_realm_id"`
	ActionType string `json:"action_type"`
	ActionId   int64  `json:"action_id"`
}
