package models

// device represents the attributes of a carbinx device.
type Device struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	Version      string `json:"version"`
	SiteID       string `json:"site_id"`
	SerialNumber string `json:"serial_number"`
	Created      string `json:"created"`
	Updated      string `json:"updated"`
}
