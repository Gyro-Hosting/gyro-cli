## gyro s3 events list-all

List all event webhooks of the service

### Synopsis

List all event webhooks of the service

Every webhook of the storage service, plus the month's usage.

API: GET /cloud/storage/{id}/events (operationId listStorageEventWebhooks, scope services_read)

```
gyro s3 events list-all <ID> [flags]
```

### Options

```
  -h, --help   help for list-all
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

* [gyro s3 events](gyro_s3_events.md)	 - List the event webhooks of a bucket

