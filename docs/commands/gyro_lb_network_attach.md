## gyro lb network attach

Attach to a private network

### Synopsis

Attach to a private network

One of your private networks in the same region (GET settings).

API: POST /cloud/load-balancers/{id}/network (operationId typed_post_load_balancers__id__network, scope services_write)

```
gyro lb network attach <ID> [flags]
```

### Options

```
      --data string              request body as JSON, @file or @- (stdin)
      --dry-run                  show the price quote / request without changing anything
  -h, --help                     help for attach
      --idempotency-key string   idempotency key (automatic for money operations)
      --network int              network
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

* [gyro lb network](gyro_lb_network.md)	 - Attach to a private network

