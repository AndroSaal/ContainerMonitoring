package entities

import (
	"time"
)

type ErrorResponse struct {
	Reason string `json:"reason"`
}

type PingInfo struct {
	IPAdress    string      `json:"ipAdress" db:"ip"`
	PingTime    time.Time   `json:"pingTime" db:"ping_time"`
	LastSuccess interface{} `json:"lastSuccessDate" db:"last_success"`
	Status      string      `json:"status" db:"status"`
}
