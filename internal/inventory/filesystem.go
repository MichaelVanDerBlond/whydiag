package inventory

import (
	"syscall"
)

type FilesystemInfo struct {
	Path        string
	TotalBytes  uint64
	FreeBytes   uint64
	UsedBytes   uint64
	UsedPercent float64
	ReadOnly    bool
}

func Filesystem(path string) (*FilesystemInfo, error) {
	var st syscall.Statfs_t

	if err := syscall.Statfs(path, &st); err != nil {
		return nil, err
	}

	total := st.Blocks * uint64(st.Bsize)
	free := st.Bavail * uint64(st.Bsize)
	used := total - free

	var usedPercent float64
	if total > 0 {
		usedPercent = float64(used) * 100 / float64(total)
	}

	// ST_RDONLY = 1 на Linux
	readOnly := st.Flags&1 != 0

	return &FilesystemInfo{
		Path:        path,
		TotalBytes:  total,
		FreeBytes:   free,
		UsedBytes:   used,
		UsedPercent: usedPercent,
		ReadOnly:    readOnly,
	}, nil
}
