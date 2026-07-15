package linux

import (
	"runtime"

	"github.com/MichaelVanDerBlond/whydiag/internal/core"
)

type Check struct{}

func New() *Check {
	return &Check{}
}

func (c *Check) Descriptor() core.Descriptor {
	return core.Descriptor{
		ID:          "linux.os",
		Name:        "Operating System",
		Category:    "linux",
		Description: "Detect operating system",
		Severity:    core.SeverityInfo,
	}
}

func (c *Check) Run(ctx *core.Context) core.Result {
	d := c.Descriptor()

	return core.Result{
		Descriptor: d,
		Status:     core.StatusOK,
		Message:    runtime.GOOS,
		Value:      runtime.GOOS,
	}
}
