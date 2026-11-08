## gyro monitor monitor resume

POST monitors:resume (GyroMonitor)

### Synopsis

POST monitors:resume (GyroMonitor)

Typed shortcut of the generic bridge: tool `monitors`, action `resume`. Body fields are the tool's own fields.

API: POST /cloud/monitors/{id}/monitors/{mid}/resume (operationId typed_post_monitors__id__monitors__mid__resume, scope services_write)

```
gyro monitor monitor resume <ID> <MID> [flags]
```

### Options

```
      --data string              request body as JSON, @file or @- (stdin)
      --dry-run                  show the price quote / request without changing anything
  -h, --help                     help for resume
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

