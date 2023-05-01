package models

// site represents the attributes of a physical location of a carbinx device.
type Site struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	Organization string `json:"organization"`
	Address      string `json:"address"`
	Postal       string `json:"postal"`
	City         string `json:"city"`
	Province     string `json:"province"`
	Country      string `json:"country"`
	Created      string `json:"created"`
	Updated      string `json:"updated"`
}
