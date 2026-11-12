## gyro s3 object retention get

Retention of an object version

### Synopsis

Retention of an object version

Lock mode, end date and legal hold of one version (`version_id` empty = current version).

API: GET /cloud/storage/{id}/buckets/{bucket}/objects/retention (operationId getStorageObjectRetention, scope services_read)

```
gyro s3 object retention get <ID> <BUCKET> [flags]
```

### Options

```
  -h, --help                help for get
      --key string          Object key
      --version-id string   Version id (empty = current)
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

