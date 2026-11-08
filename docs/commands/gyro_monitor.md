## gyro monitor

Uptime monitors

```
gyro monitor [flags]
```

### Options

```
  -h, --help   help for monitor
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

* [gyro](gyro.md)	 - Gyro Hosting command-line interface
* [gyro monitor call](gyro_monitor_call.md)	 - Call a module method of a service
* [gyro monitor channel](gyro_monitor_channel.md)	 - GET channels (GyroMonitor)
* [gyro monitor check](gyro_monitor_check.md)	 - Monitor checks
* [gyro monitor contact](gyro_monitor_contact.md)	 - Alert contacts
* [gyro monitor heartbeat](gyro_monitor_heartbeat.md)	 - GET heartbeats (GyroMonitor)
* [gyro monitor incident](gyro_monitor_incident.md)	 - Incidents
* [gyro monitor list](gyro_monitor_list.md)	 - List your services of this product line
* [gyro monitor monitor](gyro_monitor_monitor.md)	 - GET monitors (GyroMonitor)
* [gyro monitor status-page](gyro_monitor_status-page.md)	 - GET status-pages (GyroMonitor)
* [gyro monitor tool](gyro_monitor_tool.md)	 - Read a tool's data or run one of its actions (see 'tools')
* [gyro monitor tools](gyro_monitor_tools.md)	 - Show the tools, actions and safety policy of a service

