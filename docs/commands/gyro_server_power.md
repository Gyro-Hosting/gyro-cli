## gyro server power

Power on, off, shutdown, reboot or reset a server

```
gyro server power [flags]
```

### Options

```
  -h, --help   help for power
```

### Options inherited from parent commands

```
      --api-url string     API base URL (env GYRO_API_URL)
      --debug              log HTTP requests to stderr (secrets masked)
      --lang string        message language: tr or en (default from locale)
      --no-color           disable colours (also NO_COLOR)
      --no-header          omit the table header
  -o, --output string      output format: table, json, yaml, text
  -p, --profile string     profile to use (env GYRO_PROFILE)
      --query string       JMESPath query applied to the result
  -q, --quiet              print only ids
      --timeout duration   overall timeout, e.g. 30s, 10m
      --wide               show all table columns
  -y, --yes                do not ask for confirmation
```

### SEE ALSO

* [gyro server](gyro_server.md)	 - Cloud servers: create, power, snapshots, SSH ...
* [gyro server power off](gyro_server_power_off.md)	 - Hard power off
* [gyro server power on](gyro_server_power_on.md)	 - Power on
* [gyro server power reboot](gyro_server_power_reboot.md)	 - Reboot (graceful)
* [gyro server power reset](gyro_server_power_reset.md)	 - Hard reset
* [gyro server power shutdown](gyro_server_power_shutdown.md)	 - Graceful shutdown (ACPI)

