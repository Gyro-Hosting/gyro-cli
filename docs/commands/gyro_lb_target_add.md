## gyro lb target add

Add a target

### Synopsis

Add a target

type=server (server = your cloud server id, same region), type=all (all your cloud servers, new ones join automatically) or type=ip (your dedicated server IP). use_private_ip needs a private network attached.

API: POST /cloud/load-balancers/{id}/targets (operationId typed_post_load_balancers__id__targets, scope services_write)

```
gyro lb target add <ID> [flags]
```

### Options

```
      --data string              request body as JSON, @file or @- (stdin)
      --dry-run                  show the price quote / request without changing anything
  -h, --help                     help for add
      --idempotency-key string   idempotency key (automatic for money operations)
      --ip string                Dedicated server IP (type=ip)
      --server int               Cloud server id (type=server)
      --set stringArray          extra body field key=value (repeatable)
      --type string              type (server|all|ip)
      --use-private-ip           use_private_ip
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

