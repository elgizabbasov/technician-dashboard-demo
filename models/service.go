package models

// Service represents the attributes
type Service struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	DeviceID string `json:"device_id"`
	Created  string `json:"created"`
	Updated  string `json:"updated"`
}
