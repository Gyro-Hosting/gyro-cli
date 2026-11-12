## gyro s3 versions delete

Permanently delete a version

### Synopsis

Permanently delete a version

Permanently deletes one version or delete marker (cannot be undone). Versions under GOVERNANCE / COMPLIANCE retention or legal hold are refused with the end date. Needs `confirm=true`.

API: DELETE /cloud/storage/{id}/buckets/{bucket}/versions (operationId deleteStorageVersion, scope services_write)

```
gyro s3 versions delete <ID> <BUCKET> [flags]
```

### Options

```
      --dry-run                  show the price quote / request without changing anything
  -h, --help                     help for delete
      --idempotency-key string   idempotency key (automatic for money operations)
      --key string               Object key
      --version-id string        Version id
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

