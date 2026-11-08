## gyro domain update

Update domain (path: {domain})

### Synopsis

Update domain (path: {domain})

Update domain (path: {domain}). Domain operations run against the registrar module of the domain. Core WISECP client endpoint (Turkish reference: https://dev.wisecp.com/tr/sahip-olunan-alan-adlari).

API: PATCH /domains/{domain} (operationId coreUpdateDomain, scope Domains/UpdateDomain)

```
gyro domain update <DOMAIN> [flags]
```

### Options

```
      --auto-renew               Alan adı bazlı otomatik yenileme. Yalnız etkin alan adlarında ve gerçek bir mantıksal değerle.
      --billing-profile-id int   Fatura profilinin numarası. Sıfır hesabın varsayılanına döner.
      --data string              request body as JSON, @file or @- (stdin)
      --dry-run                  show the price quote / request without changing anything
  -h, --help                     help for update
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

* [gyro domain](gyro_domain.md)	 - Domains: search, register, transfer, nameservers, DNS, WHOIS

