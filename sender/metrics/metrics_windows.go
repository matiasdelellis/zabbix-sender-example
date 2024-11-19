// +build windows

package metrics

import (
	"fmt"

	"golang.org/x/sys/windows"
)

const (
	DefaultDiskPath = "c:"
)

func GetDiskSpaceAvailable(diskPath string) (float64) {
	var totalFree, callerFree, total uint64
	err := windows.GetDiskFreeSpaceEx(windows.StringToUTF16Ptr(diskPath), &callerFree, &total, &totalFree)

	if err != nil {
		fmt.Errorf("failed to call GetDiskFreeSpaceEx: %v", err)
		return -1
	}

	return float64(totalFree) / float64(total) * 100
}
