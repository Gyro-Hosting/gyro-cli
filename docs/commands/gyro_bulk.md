## gyro bulk

Run one action on up to 50 servers

### Synopsis

Run one action on up to 50 servers

API: POST /cloud/bulk (operationId bulkServerAction, scope cloud_write)

```
gyro bulk [flags]
```

### Options

```
      --action string            action (poweron|shutdown|reboot|poweroff|reset|rebuild|disable_rescue)
      --data string              request body as JSON, @file or @- (stdin)
      --dry-run                  show the price quote / request without changing anything
  -h, --help                     help for bulk
      --idempotency-key string   idempotency key (automatic for money operations)
      --image string             image
      --server-ids int64Slice    server_ids (default [])
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

* [gyro](gyro.md)	 - Gyro Hosting command-line interface

