## gyro s3 key rotate

POST access-keys:rotate (GyroDepolama)

### Synopsis

POST access-keys:rotate (GyroDepolama)

Typed shortcut of the generic bridge: tool `access-keys`, action `rotate`. Body fields are the tool's own fields. Needs `"confirm": true`.

API: POST /cloud/storage/{id}/keys/{key}/rotate (operationId typed_post_storage__id__keys__key__rotate, scope services_write)

```
gyro s3 key rotate <ID> <KEY> [flags]
```

### Options

```
      --data string              request body as JSON, @file or @- (stdin)
      --dry-run                  show the price quote / request without changing anything
  -h, --help                     help for rotate
      --idempotency-key string   idempotency key (automatic for money operations)
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

* [gyro s3 key](gyro_s3_key.md)	 - Access keys (secret shown once)

