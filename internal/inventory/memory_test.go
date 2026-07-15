package inventory

import "testing"

func TestMemory(t *testing.T) {
	mem, err := Memory()
	if err != nil {
		t.Fatal(err)
	}

	if mem.TotalKB == 0 {
		t.Fatal("total memory is zero")
	}

	if mem.AvailableKB == 0 {
		t.Fatal("available memory is zero")
	}
}
