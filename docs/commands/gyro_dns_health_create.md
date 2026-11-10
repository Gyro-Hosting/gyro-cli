## gyro dns health create

Create a health-checked record

### Synopsis

Create a health-checked record

Points `name` (A/AAAA) to `primary`; when the check fails from the majority of locations the record switches to `backup`, and back after 3 good passes. Uses one health-check add-on unit.

API: POST /cloud/dns/{id}/health-checks (operationId typed_post_dns__id__health_checks, scope services_write)

```
gyro dns health create <ID> [flags]
```

### Options

```
      --backup string            backup
      --check string             check (http|https|tcp|ping)
      --data string              request body as JSON, @file or @- (stdin)
      --dry-run                  show the price quote / request without changing anything
  -h, --help                     help for create
      --idempotency-key string   idempotency key (automatic for money operations)
      --name string              name
      --path string              path
      --port int                 port
      --primary string           primary
      --set stringArray          extra body field key=value (repeatable)
      --ttl int                  30-300 s
      --type string              type (A|AAAA)
      --zone string              zone
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

* [gyro dns health](gyro_dns_health.md)	 - List health-checked records

