package linux

import (
	"fmt"

	"github.com/MichaelVanDerBlond/whydiag/internal/core"
	"github.com/MichaelVanDerBlond/whydiag/internal/inventory"
)

type SwapCheck struct{}

func NewSwap() *SwapCheck {
	return &SwapCheck{}
}

func (s *SwapCheck) Descriptor() core.Descriptor {
	return core.Descriptor{
		ID:          "linux.swap",
		Name:        "Swap",
		Category:    "linux",
		Description: "Swap availability",
		Severity:    core.SeverityInfo,
	}
}

func (s *SwapCheck) Run(ctx *core.Context) core.Result {

	sw, err := inventory.Swap()
	if err != nil {
		return core.Result{
			Descriptor: s.Descriptor(),
			Status:     core.StatusError,
			Message:    err.Error(),
			Err:        err,
		}
	}

	status := core.StatusOK

	if sw.TotalKB == 0 {
		status = core.StatusWarning
	}

	return core.Result{
		Descriptor: s.Descriptor(),
		Status:     status,
		Message: fmt.Sprintf(
			"Total %.1f GB",
			sw.TotalGB(),
		),
		Value: sw,
	}
}
