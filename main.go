package main

import (
    "fmt"
    "os"

    "github.com/matiasdelellis/zabbix-sender-example/sender"
    "gopkg.in/ini.v1"
)


func main() {
	confPath := "/etc/zabbix_sender_example.conf"
	if _, err := os.Stat(confPath); os.IsNotExist(err) {
		confPath = "zabbix_sender_example.conf"
	}

	iniData, err := ini.Load(confPath)
	if err != nil {
		fmt.Printf("Fail to read zabbix_sender_example.conf configuration file: %v", err)
		os.Exit(1)
	}

	address := iniData.Section("Server").Key("Address").String()
	hostname := iniData.Section("Agent").Key("Hostname").String()
	every := iniData.Section("Metrics").Key("Every").MustInt(300)

	sender.Start(address, hostname, every)
}
