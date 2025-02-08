package entities

import (
	"time"
)

type ErrorResponse struct {
	Reason string `json:"reason"`
}

type PingInfo struct {
	IPAdress    string      `json:"ipAdress"`
	PingTime    time.Time   `json:"pingTime"`
	LastSuccess interface{} `json:"lastSuccessDate"`
}
