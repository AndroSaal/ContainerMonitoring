package entities

import "time"

type ErrorResponse struct {
	Reason string `json:"reason"`
}

type PingInfo struct {
	ContainerId int
	PingDate    time.Time
}
