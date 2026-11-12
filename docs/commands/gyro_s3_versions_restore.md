## gyro s3 versions restore

Restore an older version

### Synopsis

Restore an older version

Copies the version as the new current version (the old one stays). Max 5 GB; delete markers cannot be restored (use undelete).

API: POST /cloud/storage/{id}/buckets/{bucket}/versions/restore (operationId restoreStorageVersion, scope services_write)

```
gyro s3 versions restore <ID> <BUCKET> [flags]
```

### Options

```
      --data string              request body as JSON, @file or @- (stdin)
      --dry-run                  show the price quote / request without changing anything
  -h, --help                     help for restore
      --idempotency-key string   idempotency key (automatic for money operations)
      --key string               Object key
      --set stringArray          extra body field key=value (repeatable)
      --version-id string        Version id (from versions list)
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

* [gyro s3 versions](gyro_s3_versions.md)	 - List object versions

