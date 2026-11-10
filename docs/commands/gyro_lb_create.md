## gyro lb create

Order a load balancer

### Synopsis

Order a load balancer

Orders and pays from the account balance (same money rules as servers): `"confirm": true`, an `Idempotency-Key` header, `dry_run` for a quote. Ready in about a minute; then add services and targets.

API: POST /cloud/load-balancers (operationId createLoadBalancer, scope cloud_order)

```
gyro lb create [flags]
```

### Options

```
      --cycle string             cycle
      --data string              request body as JSON, @file or @- (stdin)
      --dry-run                  show the price quote / request without changing anything
      --end-customer-id int      end_customer_id
  -h, --help                     help for create
      --idempotency-key string   idempotency key (automatic for money operations)
      --product-id int           product_id
      --region string            region (eu-de-1|eu-de-2|eu-fi-1)
      --set stringArray          extra body field key=value (repeatable)
      --size string              size (small|medium|large)
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

