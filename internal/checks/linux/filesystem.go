package linux

import (
	"fmt"

	"github.com/MichaelVanDerBlond/whydiag/internal/core"
	"github.com/MichaelVanDerBlond/whydiag/internal/inventory"
)

type FilesystemCheck struct{}

func NewFilesystem() *FilesystemCheck {
	return &FilesystemCheck{}
}

func (f *FilesystemCheck) Descriptor() core.Descriptor {
	return core.Descriptor{
		ID:          "linux.filesystem",
		Name:        "Root Filesystem",
		Category:    "linux",
		Description: "Checks root filesystem health",
		Severity:    core.SeverityInfo,
	}
}

func (f *FilesystemCheck) Run(ctx *core.Context) core.Result {
	fs, err := inventory.Filesystem("/")
	if err != nil {
		return core.Result{
			Descriptor: f.Descriptor(),
			Status:     core.StatusError,
			Message:    err.Error(),
			Err:        err,
		}
	}

	status := core.StatusOK

	switch {
	case fs.ReadOnly:
		status = core.StatusError

	case fs.UsedPercent >= 95:
		status = core.StatusError

	case fs.UsedPercent >= 85:
		status = core.StatusWarning
	}

	msg := fmt.Sprintf(
		"Used %.1f%%  Free %.1f GB",
		fs.UsedPercent,
		float64(fs.FreeBytes)/1024/1024/1024,
	)

	return core.Result{
		Descriptor: f.Descriptor(),
		Status:     status,
		Message:    msg,
		Value:      fs,
	}
}
