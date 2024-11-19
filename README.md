# Zabbix Sender Example with Golang

Example of creating an small service to send custom metrics to Zabbix
as Trapper items.

## Example usage
1. Build the example with `go build` or `GOOS=windows go build`.
2. Import `zabbix_sender_example_templates.yaml` on Zabbix sever.
3. Assing `Sender Example Template` template to the Host.
4. Edit the configuration in the `zabbix_sender_example.conf` file.
5. Just launch de example with `zabbix-sender-example` or
 `zabbix-sender-example.exe`. :tada:

