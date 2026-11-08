## gyro monitor channel

GET channels (GyroMonitor)

```
gyro monitor channel [flags]
```

### Options

```
  -h, --help   help for channel
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
* [gyro monitor channel create](gyro_monitor_channel_create.md)	 - POST channels:create (GyroMonitor)
* [gyro monitor channel delete](gyro_monitor_channel_delete.md)	 - DELETE channels:delete (GyroMonitor)
* [gyro monitor channel list](gyro_monitor_channel_list.md)	 - GET channels (GyroMonitor)

