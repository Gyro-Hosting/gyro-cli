## gyro lb algorithm set

Change the balancing algorithm

### Synopsis

Change the balancing algorithm

round_robin or least_connections.

API: PUT /cloud/load-balancers/{id}/algorithm (operationId typed_put_load_balancers__id__algorithm, scope services_write)

```
gyro lb algorithm set <ID> [flags]
```

### Options

```
      --algorithm string         algorithm (round_robin|least_connections)
      --data string              request body as JSON, @file or @- (stdin)
      --dry-run                  show the price quote / request without changing anything
  -h, --help                     help for set
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

* [gyro lb algorithm](gyro_lb_algorithm.md)	 - Change the balancing algorithm

