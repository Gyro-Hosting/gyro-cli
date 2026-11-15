## gyro s3 events update

Change an event webhook

### Synopsis

Change an event webhook

Send only the fields to change. `status` active / disabled (enabling resets the failure counter). The bucket cannot change.

API: PATCH /cloud/storage/{id}/events/{event} (operationId updateStorageEventWebhook, scope services_write)

```
gyro s3 events update <ID> <EVENT> [flags]
```

### Options

```
      --data string              request body as JSON, @file or @- (stdin)
      --description string       description
      --dry-run                  show the price quote / request without changing anything
      --events strings           default: all three. Short names created / deleted / deleted_bulk are accepted
  -h, --help                     help for update
      --idempotency-key string   idempotency key (automatic for money operations)
      --prefix string            prefix
      --set stringArray          extra body field key=value (repeatable)
      --status string            status (active|disabled)
      --suffix strings           a comma separated text is accepted too
      --url string               url
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

