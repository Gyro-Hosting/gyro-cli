## gyro domain renew

Renew domain (path: {domain})

### Synopsis

Renew domain (path: {domain})

Renew domain (path: {domain}). Domain operations run against the registrar module of the domain. **Spends money / creates an invoice.** Money rules apply (AK-40): `"confirm": true` in the body, an `Idempotency-Key` header, account balance only (no stored card), daily limit of paid API operations per account (default 50, operator can override per account/reseller, resets 00:00 Europe/Istanbul); violations answer 422 with `[confirm_required]`, `[idempotency_key_required]`, `[balance_only]` or `[daily_limit]` in the message. Core WISECP client endpoint (Turkish reference: https://dev.wisecp.com/tr/sahip-olunan-alan-adlari).

API: POST /domains/{domain}/renew (operationId coreRenewDomain, scope Domains/RenewDomain)

```
gyro domain renew <DOMAIN> [flags]
```

### Options

```
      --data string              request body as JSON, @file or @- (stdin)
      --dry-run                  show the price quote / request without changing anything
  -h, --help                     help for renew
      --idempotency-key string   idempotency key (automatic for money operations)
      --set stringArray          extra body field key=value (repeatable)
      --years int                Yenileme süresi. Detaydaki fiyatlı süre satırlarından biri olmalıdır.
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

