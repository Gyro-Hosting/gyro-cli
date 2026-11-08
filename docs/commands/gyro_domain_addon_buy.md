## gyro domain addon buy

Buy domain addon (path: {domain}, {key})

### Synopsis

Buy domain addon (path: {domain}, {key})

Buy domain addon (path: {domain}, {key}). Domain operations run against the registrar module of the domain. **Spends money / creates an invoice.** Money rules apply (AK-40): `"confirm": true` in the body, an `Idempotency-Key` header, account balance only (no stored card), daily limit per account; violations answer 422 with `[confirm_required]`, `[idempotency_key_required]`, `[balance_only]` or `[daily_limit]` in the message. Core WISECP client endpoint (Turkish reference: https://dev.wisecp.com/tr/sahip-olunan-alan-adlari).

API: POST /domains/{domain}/addons/{key} (operationId coreBuyDomainAddon, scope Domains/BuyDomainAddon)

```
gyro domain addon buy <DOMAIN> <KEY> [flags]
```

### Options

```
      --dry-run                  show the price quote / request without changing anything
  -h, --help                     help for buy
      --idempotency-key string   idempotency key (automatic for money operations)
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

* [gyro domain addon](gyro_domain_addon.md)	 - Buy domain addon (path: {domain}, {key})

