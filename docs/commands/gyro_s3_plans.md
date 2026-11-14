## gyro s3 plans

List Object Storage plans

### Synopsis

List Object Storage plans

Orderable Object Storage (`standard-*`) and Protected Storage (`protected-*`) plans with limits, prices per billing cycle (fixed EUR, VAT excluded) and add-ons: extra traffic packs (`1tb`, `5tb`, one-time, this calendar month) and object lock buckets (protected plans). `available: false` = no capacity right now.

API: GET /cloud/storage/plans (operationId listStoragePlans, scope cloud_read)

```
gyro s3 plans [flags]
```

### Options

```
  -h, --help   help for plans
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

