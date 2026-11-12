## gyro s3 object-lock set

Set the default retention of a bucket

### Synopsis

Set the default retention of a bucket

Locked buckets only. `mode` GOVERNANCE or COMPLIANCE with `days` 1-3650, or OFF (new objects get no automatic retention; existing locks stay). GOVERNANCE: only Gyro staff can bypass it on request. COMPLIANCE cannot be shortened, switched to GOVERNANCE or removed by anyone (Gyro included) until it expires; the data is kept and the service keeps being billed until then, even after cancellation. Send `"confirm": true` and `"acknowledge": "COMPLIANCE"`.

API: PUT /cloud/storage/{id}/buckets/{bucket}/object-lock (operationId setStorageBucketObjectLock, scope services_write)

```
gyro s3 object-lock set <ID> <BUCKET> [flags]
```

### Options

```
      --acknowledge string       COMPLIANCE only: type COMPLIANCE to accept that it cannot be removed
      --data string              request body as JSON, @file or @- (stdin)
      --days int                 Retention days (not with OFF)
      --dry-run                  show the price quote / request without changing anything
  -h, --help                     help for set
      --idempotency-key string   idempotency key (automatic for money operations)
      --mode string              mode (GOVERNANCE|COMPLIANCE|OFF)
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

* [gyro s3 object-lock](gyro_s3_object-lock.md)	 - Default retention of a bucket

