package inventory

import (
	"os"
	"strings"
)

func KernelVersion() (string, error) {
	b, err := os.ReadFile("/proc/sys/kernel/osrelease")
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(string(b)), nil
}
