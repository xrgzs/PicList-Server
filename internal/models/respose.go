package models

type HeartbeatResponse struct {
	Success bool   `json:"success"`
	Result  string `json:"result"`
}
