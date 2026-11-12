## gyro s3 versions list

List object versions

### Synopsis

List object versions

Versions and delete markers of one folder (50 per page, `/` delimiter; sub-folders in `folders`). Continue with `key_marker` + `version_id_marker` from the previous page while `truncated` is true. In locked buckets each version carries `lock_mode`, `retain_until` and `legal_hold`.

API: GET /cloud/storage/{id}/buckets/{bucket}/versions (operationId listStorageVersions, scope services_read)

```
gyro s3 versions list <ID> <BUCKET> [flags]
```

### Options

```
  -h, --help                       help for list
      --key-marker string          Continuation: next_key_marker of the previous page
      --prefix string              Folder or key prefix (empty = root)
      --version-id-marker string   Continuation: next_version_id_marker of the previous page
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

