## gyro s3 cors set

Replace the CORS rules of a bucket

### Synopsis

Replace the CORS rules of a bucket

Replaces every rule. Send `rules` (array), `CORSRules` (an AWS `put-bucket-cors` file works as is: `gyro storage cors set ID BUCKET --data @cors.json`) or `cors_json` (the JSON as text). Field names may be AWS style (`AllowedOrigins`) or snake_case (`allowed_origins`). Limits: 100 rules, 50 origins/headers per rule, 16 KB, methods GET/PUT/POST/DELETE/HEAD, `MaxAgeSeconds` 0-86400, at most one `*` per origin. An empty list removes the rules. Rules apply to the S3 endpoint and to the static website (and custom domains).

API: PUT /cloud/storage/{id}/buckets/{bucket}/cors (operationId typed_put_storage__id__buckets__bucket__cors, scope services_write)

```
gyro s3 cors set <ID> <BUCKET> [flags]
```

### Options

```
      --cors-json string         Alternative: the rules as JSON text ({"CORSRules":[...]} or [...])
      --data string              request body as JSON, @file or @- (stdin)
      --dry-run                  show the price quote / request without changing anything
  -h, --help                     help for set
      --idempotency-key string   idempotency key (automatic for money operations)
      --rules string             CORS rules as a JSON array (JSON)
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

* [gyro s3 cors](gyro_s3_cors.md)	 - CORS rules of a bucket

