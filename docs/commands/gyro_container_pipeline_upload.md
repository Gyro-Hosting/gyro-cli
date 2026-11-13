## gyro container pipeline upload

Upload URL for an archive build

### Synopsis

Upload URL for an archive build

Returns a presigned `url` (valid 15 minutes) for one `PUT` of exactly `size` bytes (tar.gz / tgz / zip, max 200 MB; send `Content-Length: size`, e.g. `curl -T site.tar.gz "<url>"`). Then start the build with `POST .../runs {"upload_id": ...}`. Archives are deleted after 24 hours. Do not send an Idempotency-Key here (the URL would be withheld from the replayed answer).

API: POST /cloud/containers/{id}/pipelines/{pipeline}/uploads (operationId pipeline_upload, scope services_write)

```
gyro container pipeline upload <ID> <PIPELINE> [flags]
```

### Options

```
      --data string              request body as JSON, @file or @- (stdin)
      --dry-run                  show the price quote / request without changing anything
  -h, --help                     help for upload
      --idempotency-key string   idempotency key (automatic for money operations)
      --set stringArray          extra body field key=value (repeatable)
      --size int                 archive size in bytes (1 - 209715200)
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

* [gyro container pipeline](gyro_container_pipeline.md)	 - List pipelines

