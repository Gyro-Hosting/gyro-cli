## gyro s3 delete

Cancel a storage service

### Synopsis

Cancel a storage service

`confirm=true` is required. `immediate=true` cancels now: access keys, website and custom domains stop within about a minute and every bucket and object is deleted permanently within minutes (cannot be undone; Protected Storage versions under a COMPLIANCE lock or legal hold are kept until the lock ends and stay billable). Without `immediate` a cancellation request is filed and the service keeps running until it is processed.

API: DELETE /cloud/storage/{id} (operationId deleteStorageService, scope cloud_delete)

```
gyro s3 delete <ID> [flags]
```

### Options

```
      --dry-run                  show the price quote / request without changing anything
  -h, --help                     help for delete
      --idempotency-key string   idempotency key (automatic for money operations)
      --immediate                
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

* [gyro s3](gyro_s3.md)	 - S3 object storage: buckets and access keys

