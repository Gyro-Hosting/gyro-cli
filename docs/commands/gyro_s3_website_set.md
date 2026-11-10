## gyro s3 website set

Enable / update the static website

### Synopsis

Enable / update the static website

Full replace: omitted fields fall back to the defaults (index.html, no error page, SPA off). `index_document` is a file name (no `/`), served at the root and in every folder. `spa: true` answers every missing path with the entry file and status 200 (React/Vue/Angular routing, page reloads); real files are served as they are and the error page is not used while SPA is on. Enabling a disabled website needs `accept: true` (abuse rules: everything in the bucket becomes public). Address: `https://<bucket>-<service code>.web.gyro-s3.com` and connected custom domains.

API: PUT /cloud/storage/{id}/buckets/{bucket}/website (operationId typed_put_storage__id__buckets__bucket__website, scope services_write)

```
gyro s3 website set <ID> <BUCKET> [flags]
```

### Options

```
      --accept                   Accept the abuse rules (needed when the website is off)
      --data string              request body as JSON, @file or @- (stdin)
      --dry-run                  show the price quote / request without changing anything
      --error-document string    Object key of the 404 page (e.g. 404.html or errors/404.html), empty = none
  -h, --help                     help for set
      --idempotency-key string   idempotency key (automatic for money operations)
      --index-document string    Entry file name, default index.html (e.g. app.html)
      --set stringArray          extra body field key=value (repeatable)
      --spa                      Route missing sub-paths to the entry file with status 200
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

