package inventory

import "testing"

func TestHostname(t *testing.T) {
	h, err := Hostname()
	if err != nil {
		t.Fatal(err)
	}

	if h == "" {
		t.Fatal("hostname is empty")
	}
}
