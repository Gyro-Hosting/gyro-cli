## gyro s3 object retention set

Lock an object version

### Synopsis

Lock an object version

Sets or extends the retention of one version (`version_id` empty = current) for `days` 1-3650 from now. Retention can only be extended; COMPLIANCE cannot be switched to GOVERNANCE. COMPLIANCE cannot be shortened, switched to GOVERNANCE or removed by anyone (Gyro included) until it expires; the data is kept and the service keeps being billed until then, even after cancellation. Send `"confirm": true` and `"acknowledge": "COMPLIANCE"`.

API: PUT /cloud/storage/{id}/buckets/{bucket}/objects/retention (operationId setStorageObjectRetention, scope services_write)

```
gyro s3 object retention set <ID> <BUCKET> [flags]
```

### Options

```
      --acknowledge string       COMPLIANCE only: type COMPLIANCE
      --data string              request body as JSON, @file or @- (stdin)
      --days int                 days
      --dry-run                  show the price quote / request without changing anything
  -h, --help                     help for set
      --idempotency-key string   idempotency key (automatic for money operations)
      --key string               key
      --mode string              mode (GOVERNANCE|COMPLIANCE)
      --set stringArray          extra body field key=value (repeatable)
      --version-id string        Empty = current version
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

* [gyro s3 object retention](gyro_s3_object_retention.md)	 - Retention of an object version

