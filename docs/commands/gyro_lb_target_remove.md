## gyro lb target remove

Remove a target

### Synopsis

Remove a target

Same fields as add, in the query string (?type=server&server=123&confirm=true) or the JSON body. Needs `"confirm": true`.

API: DELETE /cloud/load-balancers/{id}/targets (operationId typed_delete_load_balancers__id__targets, scope services_write)

```
gyro lb target remove <ID> [flags]
```

### Options

```
      --data string              request body as JSON, @file or @- (stdin)
      --dry-run                  show the price quote / request without changing anything
  -h, --help                     help for remove
      --idempotency-key string   idempotency key (automatic for money operations)
      --ip string                Dedicated server IP
      --param-ip string          ip
      --param-server int         server
      --param-type string        type
      --server string            Cloud server id
      --set stringArray          extra body field key=value (repeatable)
      --type string              server | all | ip
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

* [gyro lb target](gyro_lb_target.md)	 - List targets with health

