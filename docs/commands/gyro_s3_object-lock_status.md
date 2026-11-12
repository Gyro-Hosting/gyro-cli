## gyro s3 object-lock status

Protected Storage: object lock summary

### Synopsis

Protected Storage: object lock summary

Locked-bucket rights (Object Lock add-on) and use, the default retention of every bucket, and the last lock scan (number of COMPLIANCE / GOVERNANCE / legal hold versions and the latest end date). Protected Storage only: a standard Object Storage service answers 409 `feature_not_available`.

API: GET /cloud/storage/{id}/object-lock (operationId getStorageObjectLock, scope services_read)

```
gyro s3 object-lock status <ID> [flags]
```

### Options

```
  -h, --help   help for status
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

