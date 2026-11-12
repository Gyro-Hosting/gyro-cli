## gyro s3 lifecycle noncurrent get

Old version expiry of a bucket

### Synopsis

Old version expiry of a bucket

Days after which non-current versions are deleted (new buckets: 30). 0 / null = kept until you delete them (they count toward the quota).

API: GET /cloud/storage/{id}/buckets/{bucket}/lifecycle/noncurrent (operationId getStorageNoncurrentExpiration, scope services_read)

```
gyro s3 lifecycle noncurrent get <ID> <BUCKET> [flags]
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

* [gyro s3 lifecycle noncurrent](gyro_s3_lifecycle_noncurrent.md)	 - Old version expiry of a bucket

