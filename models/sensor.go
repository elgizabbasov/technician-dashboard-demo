package models

// sensor represents the attributes of sensor in a carbinx device.
type Sensor struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Unit     string `json:"unit"`
	Type     string `json:"type"`
	DeviceID string `json:"device_id"`
	Created  string `json:"created"`
	Updated  string `json:"updated"`
}
