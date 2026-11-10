## gyro server app log

Live progress log of the latest app install (step changes, server log tail)

```
gyro server app log [flags]
```

### Options

```
  -h, --help   help for log
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
* [gyro server app log get](gyro_server_app_log_get.md)	 - Live progress log of the latest app install (step changes, server log tail)

