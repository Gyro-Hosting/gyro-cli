## gyro s3 bucket rename

Rename a bucket

### Synopsis

Rename a bucket

Data is not moved; the old name stops working at once (S3, website, share links). Max 5 renames per bucket per day; a released name stays reserved for 40 days. Needs `"confirm": true`.

API: POST /cloud/storage/{id}/buckets/{bucket}/rename (operationId typed_post_storage__id__buckets__bucket__rename, scope services_write)

```
gyro s3 bucket rename <ID> <BUCKET> [flags]
```

### Options

```
      --data string              request body as JSON, @file or @- (stdin)
      --dry-run                  show the price quote / request without changing anything
  -h, --help                     help for rename
      --idempotency-key string   idempotency key (automatic for money operations)
      --new-name string          new_name
      --set stringArray          extra body field key=value (repeatable)
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

* [gyro s3 bucket](gyro_s3_bucket.md)	 - Buckets

