## gyro monitor monitor

GET monitors (GyroMonitor)

```
gyro monitor monitor [flags]
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

* [gyro monitor](gyro_monitor.md)	 - Uptime monitors
* [gyro monitor monitor create](gyro_monitor_monitor_create.md)	 - POST monitors:create (GyroMonitor)
* [gyro monitor monitor delete](gyro_monitor_monitor_delete.md)	 - DELETE monitors:delete (GyroMonitor)
* [gyro monitor monitor list](gyro_monitor_monitor_list.md)	 - GET monitors (GyroMonitor)
* [gyro monitor monitor pause](gyro_monitor_monitor_pause.md)	 - POST monitors:pause (GyroMonitor)
* [gyro monitor monitor resume](gyro_monitor_monitor_resume.md)	 - POST monitors:resume (GyroMonitor)
* [gyro monitor monitor update](gyro_monitor_monitor_update.md)	 - PATCH monitors:edit (GyroMonitor)

