package linux

import (
	"github.com/MichaelVanDerBlond/whydiag/internal/core"
	"github.com/MichaelVanDerBlond/whydiag/internal/inventory"
)

type KernelCheck struct{}

func NewKernel() *KernelCheck {
	return &KernelCheck{}
}

func (k *KernelCheck) Descriptor() core.Descriptor {
	return core.Descriptor{
		ID:          "linux.kernel",
		Name:        "Kernel",
		Category:    "linux",
		Description: "Linux kernel version",
		Severity:    core.SeverityInfo,
	}
}

func (k *KernelCheck) Run(ctx *core.Context) core.Result {

	v, err := inventory.KernelVersion()
	if err != nil {
		return core.Result{
			Descriptor: k.Descriptor(),
			Status:     core.StatusError,
			Message:    err.Error(),
			Err:        err,
		}
	}

	return core.Result{
		Descriptor: k.Descriptor(),
		Status:     core.StatusOK,
		Message:    v,
		Value:      v,
	}
}
