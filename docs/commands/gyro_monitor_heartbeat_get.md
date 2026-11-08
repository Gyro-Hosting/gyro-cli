## gyro monitor heartbeat get

GET heartbeats (GyroMonitor)

### Synopsis

GET heartbeats (GyroMonitor)

Typed shortcut of the generic bridge: tool `heartbeats`. Body fields are the tool's own fields.

API: GET /cloud/monitors/{id}/heartbeats (operationId typed_get_monitors__id__heartbeats, scope services_read)

```
gyro monitor heartbeat get <ID> [flags]
```

### Options

```
  -h, --help   help for get
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

* [gyro monitor heartbeat](gyro_monitor_heartbeat.md)	 - GET heartbeats (GyroMonitor)

