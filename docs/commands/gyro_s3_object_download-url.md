## gyro s3 object download-url

Presigned download URL

### Synopsis

Presigned download URL

Returns a 15-minute GET URL (Content-Disposition: attachment). Downloads count toward the service's monthly traffic.

API: POST /cloud/storage/{id}/buckets/{bucket}/objects/download-url (operationId typed_post_storage__id__buckets__bucket__objects_download_url, scope services_write)

```
gyro s3 object download-url <ID> <BUCKET> [flags]
```

### Options

```
      --data string              request body as JSON, @file or @- (stdin)
      --dry-run                  show the price quote / request without changing anything
  -h, --help                     help for download-url
      --idempotency-key string   idempotency key (automatic for money operations)
      --key string               Object key
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

* [gyro s3 object](gyro_s3_object.md)	 - List objects in a bucket folder

