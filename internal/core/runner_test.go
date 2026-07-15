package core

import "testing"

type testCheck struct{}

func (t testCheck) Descriptor() Descriptor {
	return Descriptor{
		ID:       "test.check",
		Name:     "test",
		Category: "test",
		Severity: SeverityInfo,
	}
}

func (t testCheck) Run(ctx *Context) Result {
	return Result{
		Descriptor: t.Descriptor(),
		Status:     StatusOK,
		Message:    "ok",
		Value:      "ok",
	}
}

func TestRunnerRegister(t *testing.T) {
	r := NewRunner()

	if len(r.checks) != 0 {
		t.Fatalf("expected 0 checks")
	}

	r.Register(testCheck{})

	if len(r.checks) != 1 {
		t.Fatalf("expected 1 check")
	}
}

func TestCheckRun(t *testing.T) {
	c := testCheck{}

	res := c.Run(NewContext())

	if res.Status != StatusOK {
		t.Fatal("unexpected status")
	}

	if res.Value != "ok" {
		t.Fatal("unexpected value")
	}

	if res.Descriptor.ID != "test.check" {
		t.Fatal("unexpected descriptor")
	}
}
