package entities

import (
	"time"
)

type ErrorResponse struct {
	Reason string `json:"reason"`
}

type PingInfo struct {
	IPAdress    string    `json:"ipAdress" db:"ip"`
	PingTime    time.Time `json:"pingTime" db:"ping_time"`
	LastSuccess time.Time `json:"lastSuccessDate" db:"last_success"`
	Status      string    `json:"status" db:"status"`
}

type PingInfoResponse struct {
	IPAdress    string `json:"ipAdress"`
	PingTime    string `json:"pingTime"`
	LastSuccess string `json:"lastSuccessDate"`
	Status      string `json:"status"`
}
