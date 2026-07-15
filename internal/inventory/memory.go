package inventory

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type MemoryInfo struct {
	TotalKB     uint64
	AvailableKB uint64
}

func Memory() (*MemoryInfo, error) {
	f, err := os.Open("/proc/meminfo")
	if err != nil {
		return nil, err
	}
	defer f.Close()

	info := &MemoryInfo{}

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()

		switch {
		case strings.HasPrefix(line, "MemTotal:"):
			v, err := parseKB(line)
			if err != nil {
				return nil, err
			}
			info.TotalKB = v

		case strings.HasPrefix(line, "MemAvailable:"):
			v, err := parseKB(line)
			if err != nil {
				return nil, err
			}
			info.AvailableKB = v
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return info, nil
}

func parseKB(line string) (uint64, error) {
	fields := strings.Fields(line)

	if len(fields) < 2 {
		return 0, fmt.Errorf("invalid meminfo line: %q", line)
	}

	return strconv.ParseUint(fields[1], 10, 64)
}

func (m *MemoryInfo) TotalGB() float64 {
	return float64(m.TotalKB) / 1024 / 1024
}

func (m *MemoryInfo) AvailableGB() float64 {
	return float64(m.AvailableKB) / 1024 / 1024
}
