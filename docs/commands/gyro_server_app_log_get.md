## gyro server app log get

Live progress log of the latest app install (step changes, server log tail)

### Synopsis

Live progress log of the latest app install (step changes, server log tail)

Lines with `id` greater than `after`, oldest first (poll with the last id you have). `source` is step, server, system or staff. Server lines are untrusted text sent by the server (secrets masked); render them as plain text.

API: GET /cloud/servers/{id}/app/log (operationId getAppInstallLog, scope cloud_read)

```
gyro server app log get <ID> [flags]
```

### Options

```
      --after int         
  -h, --help              help for get
      --param-limit int   
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

* [gyro server app log](gyro_server_app_log.md)	 - Live progress log of the latest app install (step changes, server log tail)

