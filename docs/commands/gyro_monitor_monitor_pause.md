## gyro monitor monitor pause

POST monitors:pause (GyroMonitor)

### Synopsis

POST monitors:pause (GyroMonitor)

Typed shortcut of the generic bridge: tool `monitors`, action `pause`. Body fields are the tool's own fields.

API: POST /cloud/monitors/{id}/monitors/{mid}/pause (operationId typed_post_monitors__id__monitors__mid__pause, scope services_write)

```
gyro monitor monitor pause <ID> <MID> [flags]
```

### Options

```
      --data string              request body as JSON, @file or @- (stdin)
      --dry-run                  show the price quote / request without changing anything
  -h, --help                     help for pause
      --idempotency-key string   idempotency key (automatic for money operations)
      --set stringArray          extra body field key=value (repeatable)
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

* [gyro monitor monitor](gyro_monitor_monitor.md)	 - GET monitors (GyroMonitor)

