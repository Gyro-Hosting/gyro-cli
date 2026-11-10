## gyro lb types

List load balancer sizes and prices

### Synopsis

List load balancer sizes and prices

Orderable sizes (small/medium/large) with limits, regions and EUR prices per cycle (excl. VAT); `list_amount` is the strike-through list price.

API: GET /cloud/load-balancer-types (operationId listLoadBalancerTypes, scope cloud_read)

```
gyro lb types [flags]
```

### Options

```
  -h, --help   help for types
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

* [gyro lb](gyro_lb.md)	 - List load balancers

