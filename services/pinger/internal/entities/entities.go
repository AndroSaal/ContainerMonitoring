package entities

import (
	"time"
)

type ErrorResponse struct {
	Reason string `json:"reason"`
}

type PingInfo struct {
	IPAddress   string      `json:"ipAdress"`
	PingTime    time.Time   `json:"pingTime"`
	LastSuccess interface{} `json:"lastSuccessDate"`
	Status      string      `json:"status"`
}

type Ips struct {
	Ips []string `json:"ips"`
}
