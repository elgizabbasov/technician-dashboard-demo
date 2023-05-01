package models

// sensor represents the attributes of sensor in a carbinx device.
type SensorData struct {
	ID       string `json:"id"`
	SensorID string `json:"sensor_id"`
	Data     string `json:"data"`
	Created  string `json:"created"`
}
