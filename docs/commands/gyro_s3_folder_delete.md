## gyro s3 folder delete

Delete a folder with its contents

### Synopsis

Delete a folder with its contents

Deletes every object under `prefix` (query). Large folders are deleted in batches: repeat while `more` is true. Needs `"confirm": true`.

API: DELETE /cloud/storage/{id}/buckets/{bucket}/folders (operationId typed_delete_storage__id__buckets__bucket__folders, scope services_write)

```
gyro s3 folder delete <ID> <BUCKET> [flags]
```

### Options

```
      --data string              request body as JSON, @file or @- (stdin)
      --dry-run                  show the price quote / request without changing anything
  -h, --help                     help for delete
      --idempotency-key string   idempotency key (automatic for money operations)
      --prefix string            Folder, ends with /
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

* [gyro s3 folder](gyro_s3_folder.md)	 - Count objects in a folder

