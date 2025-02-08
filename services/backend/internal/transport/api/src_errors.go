package api

import "errors"

var (
	ErrNotFound        = errors.New("not found")
	ErrNotFoundMassage = "not found"

	ErrServerDown = "The server is currently overloaded, please try again later"
)
