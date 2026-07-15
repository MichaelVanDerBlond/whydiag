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

	for _, c := range r.checks {
		start := time.Now()

		res := c.Run(ctx)

		res.Duration = time.Since(start)

		if res.Descriptor.Name == "" {
			res.Descriptor = c.Descriptor()
		}

		fmt.Printf(
			"%-25s %-8s %s\n",
			res.Descriptor.Name,
			res.Status,
			res.Message,
		)
	}
}
