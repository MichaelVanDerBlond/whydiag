package linux

import (
	"fmt"

	"github.com/MichaelVanDerBlond/whydiag/internal/core"
	"github.com/MichaelVanDerBlond/whydiag/internal/inventory"
)

type MemoryCheck struct{}

func NewMemory() *MemoryCheck {
	return &MemoryCheck{}
}

func (m *MemoryCheck) Descriptor() core.Descriptor {
	return core.Descriptor{
		ID:          "linux.memory",
		Name:        "System Memory",
		Category:    "linux",
		Description: "Checks available system memory",
		Severity:    core.SeverityInfo,
	}
}

func (m *MemoryCheck) Run(ctx *core.Context) core.Result {
	mem, err := inventory.Memory()
	if err != nil {
		return core.Result{
			Descriptor: m.Descriptor(),
			Status:     core.StatusError,
			Message:    err.Error(),
			Err:        err,
		}
	}

	total := mem.TotalGB()
	avail := mem.AvailableGB()

	status := core.StatusOK
	message := fmt.Sprintf(
		"Total %.1f GB, Available %.1f GB",
		total,
		avail,
	)

	switch {
	case total < 2:
		status = core.StatusError
	case total < 4:
		status = core.StatusWarning
	}

	return core.Result{
		Descriptor: m.Descriptor(),
		Status:     status,
		Message:    message,
		Value:      mem,
	}
}
