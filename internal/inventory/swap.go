package inventory

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type SwapInfo struct {
	TotalKB uint64
	FreeKB  uint64
}

func Swap() (*SwapInfo, error) {

	f, err := os.Open("/proc/meminfo")
	if err != nil {
		return nil, err
	}
	defer f.Close()

	info := &SwapInfo{}

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {

		fields := strings.Fields(scanner.Text())

		if len(fields) < 2 {
			continue
		}

		switch fields[0] {

		case "SwapTotal:":
			v, err := strconv.ParseUint(fields[1], 10, 64)
			if err != nil {
				return nil, err
			}
			info.TotalKB = v

		case "SwapFree:":
			v, err := strconv.ParseUint(fields[1], 10, 64)
			if err != nil {
				return nil, err
			}
			info.FreeKB = v

		}

	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return info, nil
}

func (s *SwapInfo) TotalGB() float64 {
	return float64(s.TotalKB) / 1024 / 1024
}

func (s *SwapInfo) FreeGB() float64 {
	return float64(s.FreeKB) / 1024 / 1024
}

func (s *SwapInfo) String() string {
	return fmt.Sprintf("%.1f GB", s.TotalGB())
}
