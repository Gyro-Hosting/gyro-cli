## gyro s3 object legal-hold get

Legal hold of an object version

### Synopsis

Legal hold of an object version

`legal_hold` true = the version cannot be deleted until the hold is removed (no end date).

API: GET /cloud/storage/{id}/buckets/{bucket}/objects/legal-hold (operationId getStorageObjectLegalHold, scope services_read)

```
gyro s3 object legal-hold get <ID> <BUCKET> [flags]
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

* [gyro s3 object legal-hold](gyro_s3_object_legal-hold.md)	 - Legal hold of an object version

