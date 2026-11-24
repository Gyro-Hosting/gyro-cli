## gyro s3 events deliveries

Recent deliveries of a webhook

### Synopsis

Recent deliveries of a webhook

The last 20 delivery records (sent / failed / retrying, HTTP code, error, attempts), newest first; kept 30 days.

API: GET /cloud/storage/{id}/events/{event}/deliveries (operationId listStorageEventDeliveries, scope services_read)

```
gyro s3 events deliveries <ID> <EVENT> [flags]
```

### Options

```
  -h, --help   help for deliveries
```

### Options inherited from parent commands

```
      --account gyro account list   act in another account you are a panel user of (id from gyro account list; "self" = your own; env GYRO_ACCOUNT, profile default: `gyro account use`)
      --api-url string              API base URL (env GYRO_API_URL)
      --debug                       log HTTP requests to stderr (secrets masked)
      --lang string                 message language: tr or en (default from locale)
      --no-color                    disable colours (also NO_COLOR)
      --no-header                   omit the table header
  -o, --output string               output format: table, json, yaml, text
  -p, --profile string              profile to use (env GYRO_PROFILE)
      --query string                JMESPath query applied to the result
  -q, --quiet                       print only ids
      --timeout duration            overall timeout, e.g. 30s, 10m
      --wide                        show all table columns
  -y, --yes                         do not ask for confirmation
```

### SEE ALSO

* [gyro s3 events](gyro_s3_events.md)	 - List the event webhooks of a bucket

