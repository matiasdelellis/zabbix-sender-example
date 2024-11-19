package sender

import (
	"fmt"
	"time"

	"github.com/matiasdelellis/zabbix-sender-example/sender/metrics"

	"github.com/essentialkaos/go-zabbix"
)


func Start(address string, hostname string, every int) {
	client, err := zabbix.NewClient(address, hostname)

	if err != nil {
		fmt.Printf("zabbix.NewClient error: %v \n: ", err)
		return
	}

	fmt.Printf("conection successfully to: %v \n", address)

	for {
		osHostname := metrics.GetOsHostname()
		diskPath := metrics.GetDefaultDiskPath()
		spaceAvailable := metrics.GetDiskSpaceAvailable(diskPath)

		fmt.Printf("Hostname: %v \n", osHostname)
		fmt.Printf("Disk path: %v \n", diskPath)
		fmt.Printf("Disk space available: %.2f %% \n", spaceAvailable)

		client.Add("example.hostname", osHostname)
		client.Add("example.disk_path", diskPath)
		client.Add("example.disk_space_available", spaceAvailable)

		res, err := client.Send()
		if err != nil {
			fmt.Printf("Send error: %v \n", err)
		}

		fmt.Printf("Metrics sended (total %d | processed: %d | failed: %d) \n", res.Total, res.Processed, res.Failed)

		time.Sleep(time.Duration(every) * time.Second)
	}
}
