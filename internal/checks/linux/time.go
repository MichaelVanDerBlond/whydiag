package linux

import (
	"github.com/MichaelVanDerBlond/whydiag/internal/core"
	"github.com/MichaelVanDerBlond/whydiag/internal/inventory"
)

type TimeCheck struct{}

func NewTime() *TimeCheck {
	return &TimeCheck{}
}

func (t *TimeCheck) Descriptor() core.Descriptor {
	return core.Descriptor{
		ID:          "linux.timezone",
		Name:        "Timezone",
		Category:    "linux",
		Description: "Current timezone",
		Severity:    core.SeverityInfo,
	}
}

func (t *TimeCheck) Run(ctx *core.Context) core.Result {

	tz := inventory.Timezone()

	return core.Result{
		Descriptor: t.Descriptor(),
		Status:     core.StatusOK,
		Message:    tz,
		Value:      tz,
	}
}
