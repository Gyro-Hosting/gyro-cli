## gyro s3 object upload-url

Presigned upload URLs

### Synopsis

Presigned upload URLs

Up to 100 files per call, max 5 GiB each. PUT the file body to each `url` within 15 minutes; data never passes through the API. Larger files: any S3 client (multipart).

API: POST /cloud/storage/{id}/buckets/{bucket}/objects/upload-url (operationId typed_post_storage__id__buckets__bucket__objects_upload_url, scope services_write)

```
gyro s3 object upload-url <ID> <BUCKET> [flags]
```

### Options

```
      --data string              request body as JSON, @file or @- (stdin)
      --dry-run                  show the price quote / request without changing anything
  -h, --help                     help for upload-url
      --idempotency-key string   idempotency key (automatic for money operations)
      --items strings            items
      --prefix string            Target folder (ends with /)
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

