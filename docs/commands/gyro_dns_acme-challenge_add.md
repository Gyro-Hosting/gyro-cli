## gyro dns acme-challenge add

Add an ACME challenge TXT value

### Synopsis

Add an ACME challenge TXT value

Adds `value` to the TXT record set at `name` (TTL 60); other values stay, so the apex and the wildcard of one certificate can be validated together (max 10 values per name). The service and the zone are found from the name (longest Gyro DNS zone of yours that covers it; hosting/cPanel zones are not supported). Idempotent.

API: POST /cloud/dns/acme-challenge (operationId acmeChallengeAdd, scope dns_acme)

```
gyro dns acme-challenge add [flags]
```

### Options

```
      --data string              request body as JSON, @file or @- (stdin)
      --dry-run                  show the price quote / request without changing anything
  -h, --help                     help for add
      --idempotency-key string   idempotency key (automatic for money operations)
      --name string              _acme-challenge.<domain> (for *.example.com: _acme-challenge.example.com)
      --set stringArray          extra body field key=value (repeatable)
      --value string             Challenge value (base64url, 16-128 characters)
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

* [gyro dns acme-challenge](gyro_dns_acme-challenge.md)	 - Add an ACME challenge TXT value

