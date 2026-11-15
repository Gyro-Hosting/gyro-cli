## gyro s3 events test

Send a test event

### Synopsis

Send a test event

Sends a `storage.test` event to the webhook now and waits up to 15 s for the answer (not counted against the allowance; 10 tests per 10 minutes).

API: POST /cloud/storage/{id}/events/{event}/test (operationId testStorageEventWebhook, scope services_write)

```
gyro s3 events test <ID> <EVENT> [flags]
```

### Options

```
      --dry-run                  show the price quote / request without changing anything
  -h, --help                     help for test
      --idempotency-key string   idempotency key (automatic for money operations)
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

