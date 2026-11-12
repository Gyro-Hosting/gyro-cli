## gyro s3 object legal-hold set

Place or remove a legal hold

### Synopsis

Place or remove a legal hold

`status` on / off (locked buckets only). A legal hold blocks deletion with no end date, independent of retention; you can remove it yourself.

API: PUT /cloud/storage/{id}/buckets/{bucket}/objects/legal-hold (operationId setStorageObjectLegalHold, scope services_write)

```
gyro s3 object legal-hold set <ID> <BUCKET> [flags]
```

### Options

```
      --data string              request body as JSON, @file or @- (stdin)
      --dry-run                  show the price quote / request without changing anything
  -h, --help                     help for set
      --idempotency-key string   idempotency key (automatic for money operations)
      --key string               key
      --set stringArray          extra body field key=value (repeatable)
      --status string            status (on|off)
      --version-id string        Empty = current version
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

