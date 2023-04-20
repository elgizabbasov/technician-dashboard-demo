package models

// Service data is used to store Report Card info from technicians.
type ServiceData struct {
	ID             string   `json:"id"`
	Data           string   `json:"data"`
	ServiceID      string   `json:"service_id"`
	Created        string   `json:"created"`
	Updated        string   `json:"updated"`
	Pearlash       []string `json:"pearlash"`
	Mechanicalroom []string `json:"mechanicalroom"`
	Pails          []string `json:"pails"`
}
