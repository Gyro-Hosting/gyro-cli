## gyro service renew

Renew service (path: {id})

### Synopsis

Renew service (path: {id})

Renew service (path: {id}). Service of the calling account (any product line). **Spends money / creates an invoice.** Money rules apply (AK-40): `"confirm": true` in the body, an `Idempotency-Key` header, account balance only (no stored card), daily limit of paid API operations per account (default 50, operator can override per account/reseller, resets 00:00 Europe/Istanbul); violations answer 422 with `[confirm_required]`, `[idempotency_key_required]`, `[balance_only]` or `[daily_limit]` in the message. Core WISECP client endpoint (Turkish reference: https://dev.wisecp.com/tr/sahip-olunan-hizmetler).

API: POST /services/{id}/renew (operationId coreRenewService, scope Services/RenewService)

```
gyro service renew <ID> [flags]
```

### Options

```
      --cycle string             Yenileme döngüsü. Hizmet detayındaki döngü listesinden seçilir; verilmezse mevcut döngü kullanılır.
      --data string              request body as JSON, @file or @- (stdin)
      --dry-run                  show the price quote / request without changing anything
  -h, --help                     help for renew
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

* [gyro service](gyro_service.md)	 - All services of every product line (generic bridge)

