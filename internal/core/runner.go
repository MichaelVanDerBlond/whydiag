package core

import (
	"fmt"
	"time"
)

type Runner struct {
	checks []Check
}

func NewRunner() *Runner {
	return &Runner{}
}

func (r *Runner) Register(c Check) {
	r.checks = append(r.checks, c)
}

func (r *Runner) Run() {
	ctx := NewContext()

	var okCount, warnCount, errCount, skipCount int

	for _, c := range r.checks {
		start := time.Now()

		res := c.Run(ctx)
		res.Duration = time.Since(start)

		if res.Descriptor.Name == "" {
			res.Descriptor = c.Descriptor()
		}

		switch res.Status {
		case StatusOK:
			okCount++
		case StatusWarning:
			warnCount++
		case StatusError:
			errCount++
		case StatusSkip:
			skipCount++
		}

		fmt.Printf(
			"[%7s] %-28s %s\n",
			res.Status,
			res.Descriptor.Name,
			res.Message,
		)
	}

	fmt.Println()
	fmt.Println("Summary")
	fmt.Println("-------")
	fmt.Printf("OK      : %d\n", okCount)
	fmt.Printf("WARNING : %d\n", warnCount)
	fmt.Printf("ERROR   : %d\n", errCount)
	fmt.Printf("SKIP    : %d\n", skipCount)
}
