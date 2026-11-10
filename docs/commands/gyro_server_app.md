## gyro server app

App install status (credentials shown once when ready)

```
gyro server app [flags]
```

### Options

```
  -h, --help   help for app
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
* [gyro server app create](gyro_server_app_create.md)	 - Install an app (rebuilds the disk; confirm required; 3 per server per day)
* [gyro server app get](gyro_server_app_get.md)	 - App install status (credentials shown once when ready)
* [gyro server app log](gyro_server_app_log.md)	 - Live progress log of the latest app install (step changes, server log tail)
* [gyro server app retry](gyro_server_app_retry.md)	 - Retry a failed or silent app install (rebuilds the disk; confirm required; 3 per server per hour)

