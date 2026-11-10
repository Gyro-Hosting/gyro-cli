## gyro server app retry

Retry a failed or silent app install (rebuilds the disk; confirm required; 3 per server per hour)

### Synopsis

Retry a failed or silent app install (rebuilds the disk; confirm required; 3 per server per hour)

Allowed when the latest install has `status` hata (failed) or `stuck` = true (no news for too long). On cloud servers the server is rebuilt with the same app, OS and domain; ALL data on the disk is deleted.

API: POST /cloud/servers/{id}/app/retry (operationId retryAppInstall, scope cloud_write)

```
gyro server app retry <ID> [flags]
```

### Options

```
      --data string              request body as JSON, @file or @- (stdin)
      --dry-run                  show the price quote / request without changing anything
  -h, --help                     help for retry
      --idempotency-key string   idempotency key (automatic for money operations)
      --set stringArray          extra body field key=value (repeatable)
      --wait                     wait until the server is no longer busy
      --wait-timeout duration    maximum time for --wait (default 15m0s)
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

