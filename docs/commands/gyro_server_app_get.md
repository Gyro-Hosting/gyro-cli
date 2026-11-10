## gyro server app get

App install status (credentials shown once when ready)

### Synopsis

App install status (credentials shown once when ready)

Since 2.7.0 also `stuck` (open but silent for too long), `stuck_reason` (`no_callback` = the server never reported back), `last_update_at`, `attempt`, `can_retry`, `error_code`, `step_text`. Live log: GET /cloud/servers/{id}/app/log.

API: GET /cloud/servers/{id}/app (operationId getAppInstall, scope cloud_read)

```
gyro server app get <ID> [flags]
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

* [gyro server app](gyro_server_app.md)	 - App install status (credentials shown once when ready)

