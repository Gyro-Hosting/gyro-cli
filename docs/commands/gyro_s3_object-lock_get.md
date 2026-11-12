## gyro s3 object-lock get

Default retention of a bucket

### Synopsis

Default retention of a bucket

`object_lock` (set at creation, cannot change) and `default_retention` (null = none): new objects get this retention automatically.

API: GET /cloud/storage/{id}/buckets/{bucket}/object-lock (operationId getStorageBucketObjectLock, scope services_read)

```
gyro s3 object-lock get <ID> <BUCKET> [flags]
```

### Options

```
  -h, --help   help for get
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

