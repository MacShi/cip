package model
type ClientINfo struct {
	IpAddress string `json:"ip_address"`
	UserAgent string `json:"user_agent"`
	Location Location `json:"location"`
}
type Location struct {
	Province string `json:"province"`
	City	 string `json:"city"`
}
