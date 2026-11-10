## gyro s3 website get

Static website settings of a bucket

### Synopsis

Static website settings of a bucket

`enabled`, `url`, `index_document` (entry file, default index.html), `error_document` (404 page) and `spa` (single-page app fallback).

API: GET /cloud/storage/{id}/buckets/{bucket}/website (operationId typed_get_storage__id__buckets__bucket__website, scope services_read)

```
gyro s3 website get <ID> <BUCKET> [flags]
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

* [gyro s3 website](gyro_s3_website.md)	 - Static website settings of a bucket

