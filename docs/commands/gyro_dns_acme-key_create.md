## gyro dns acme-key create

Create an ACME key

### Synopsis

Create an ACME key

A key for certificate automation that can only add/remove `_acme-challenge` TXT records (scope `dns_acme`), optionally only in `zones` and only from `ips`. It cannot read zones, change other records or reach any other endpoint. The token is returned once. Max 10 per service.

API: POST /cloud/dns/{id}/acme-keys (operationId createAcmeKey, scope services_write)

```
gyro dns acme-key create <ID> [flags]
```

### Options

```
      --data string              request body as JSON, @file or @- (stdin)
      --dry-run                  show the price quote / request without changing anything
  -h, --help                     help for create
      --idempotency-key string   idempotency key (automatic for money operations)
      --ips strings              Allowed client IPs (empty = any)
      --name string              Label (e.g. web-01)
      --set stringArray          extra body field key=value (repeatable)
      --zones strings            Allowed zones (empty = all zones of the service)
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

* [gyro dns acme-key](gyro_dns_acme-key.md)	 - List ACME keys

