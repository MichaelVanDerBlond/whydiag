package linux

import (
	"github.com/MichaelVanDerBlond/whydiag/internal/core"
	"github.com/MichaelVanDerBlond/whydiag/internal/inventory"
)

type HostnameCheck struct{}

func NewHostname() *HostnameCheck {
	return &HostnameCheck{}
}

func (h *HostnameCheck) Descriptor() core.Descriptor {
	return core.Descriptor{
		ID:          "linux.hostname",
		Name:        "Hostname",
		Category:    "linux",
		Description: "Checks system hostname",
		Severity:    core.SeverityInfo,
	}
}

func (h *HostnameCheck) Run(ctx *core.Context) core.Result {
	name, err := inventory.Hostname()
	if err != nil {
		return core.Result{
			Descriptor: h.Descriptor(),
			Status:     core.StatusError,
			Message:    err.Error(),
			Err:        err,
		}
	}

	status := core.StatusOK

	if name == "localhost" || name == "localhost.localdomain" {
		status = core.StatusWarning
	}

	return core.Result{
		Descriptor: h.Descriptor(),
		Status:     status,
		Message:    name,
		Value:      name,
	}
}
