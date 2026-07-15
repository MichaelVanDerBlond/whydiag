package inventory

import "testing"

func TestFilesystem(t *testing.T) {
	fs, err := Filesystem("/")
	if err != nil {
		t.Fatal(err)
	}

	if fs.TotalBytes == 0 {
		t.Fatal("filesystem size is zero")
	}
}
