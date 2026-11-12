## gyro s3 object-lock scan

Scan locked versions now

### Synopsis

Scan locked versions now

Counts locked versions in every locked bucket (up to 25 s / 5000 versions; `complete: false` when cut short).

API: POST /cloud/storage/{id}/object-lock/scan (operationId scanStorageObjectLock, scope services_write)

```
gyro s3 object-lock scan <ID> [flags]
```

### Options

```
      --dry-run                  show the price quote / request without changing anything
  -h, --help                     help for scan
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

* [gyro s3 object-lock](gyro_s3_object-lock.md)	 - Default retention of a bucket

