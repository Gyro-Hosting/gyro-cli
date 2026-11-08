## gyro monitor check

Monitor checks

```
gyro monitor check [flags]
```

### Options

```
  -h, --help   help for check
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
* [gyro monitor check create](gyro_monitor_check_create.md)	 - create (fields with --set key=value)
* [gyro monitor check delete](gyro_monitor_check_delete.md)	 - delete (fields with --set key=value)
* [gyro monitor check list](gyro_monitor_check_list.md)	 - Show
* [gyro monitor check update](gyro_monitor_check_update.md)	 - update (fields with --set key=value)

