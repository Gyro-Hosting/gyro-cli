## gyro s3 lifecycle noncurrent set

Set old version expiry

### Synopsis

Set old version expiry

`days` 1-3650, or 0 to keep old versions. Locked versions are deleted only after their retention ends. Other lifecycle rules of the bucket are kept.

API: PUT /cloud/storage/{id}/buckets/{bucket}/lifecycle/noncurrent (operationId setStorageNoncurrentExpiration, scope services_write)

```
gyro s3 lifecycle noncurrent set <ID> <BUCKET> [flags]
```

### Options

```
      --data string              request body as JSON, @file or @- (stdin)
      --days int                 days
      --dry-run                  show the price quote / request without changing anything
  -h, --help                     help for set
      --idempotency-key string   idempotency key (automatic for money operations)
      --set stringArray          extra body field key=value (repeatable)
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

* [gyro s3 lifecycle noncurrent](gyro_s3_lifecycle_noncurrent.md)	 - Old version expiry of a bucket

