package metrics

import (
	"fmt"
	"os"
)

func GetOsHostname() (string) {
	osHostname, err := os.Hostname()
	if err != nil {
		fmt.Printf("failed to call os.Hostname: %v", err)
		return "unknown"
	}
	return osHostname
}

func GetDefaultDiskPath() (string) {
	return DefaultDiskPath
}