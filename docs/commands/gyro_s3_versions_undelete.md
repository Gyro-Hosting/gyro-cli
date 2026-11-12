## gyro s3 versions undelete

Remove a delete marker

### Synopsis

Remove a delete marker

Removes the latest delete marker of the key, so the previous version becomes current again. `version_id` must be that delete marker.

API: POST /cloud/storage/{id}/buckets/{bucket}/versions/undelete (operationId undeleteStorageObject, scope services_write)

```
gyro s3 versions undelete <ID> <BUCKET> [flags]
```

### Options

```
      --data string              request body as JSON, @file or @- (stdin)
      --dry-run                  show the price quote / request without changing anything
  -h, --help                     help for undelete
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

