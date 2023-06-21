package domain

type Device struct {
	Ipv4       string `json:"ipv4"`
	MacAddress string `json:"mac_address"`
	ModelID    string `json:"model_id"`
	Name       string `json:"name"`
	SampleRate int    `json:"sample_rate"`
	ServerName string `json:"server_name"`
}
