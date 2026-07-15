package core

import "time"

type Status string

const (
	StatusOK      Status = "OK"
	StatusWarning Status = "WARNING"
	StatusError   Status = "ERROR"
	StatusSkip    Status = "SKIP"
)

type Result struct {
	Descriptor Descriptor

	Status Status

	Message string

	Value any

	Duration time.Duration

	Err error
}
