// +build !windows

package metrics

import (
	"fmt"

	"golang.org/x/sys/unix"
)

const (
	DefaultDiskPath = "/"
)

func GetDiskSpaceAvailable(diskPath string) (float64) {
	var stat unix.Statfs_t

	err := unix.Statfs(diskPath, &stat)

	if err != nil {
		fmt.Printf("failed to call unix.Statfs: %v", err)
		return -1.0
	}

	return float64(stat.Bfree) / float64(stat.Blocks) * 100
}
