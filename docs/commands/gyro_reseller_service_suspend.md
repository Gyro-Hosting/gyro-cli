## gyro reseller service suspend

Suspend a service (reseller)

### Synopsis

Suspend a service (reseller)

Suspends an active service (e.g. the end-customer did not pay the reseller). Data is kept. Repeating on a suspended service returns `changed: false`. 422 `invalid_service_state` when the service is not active. Only services on the reseller's own account (404 `service_not_found` otherwise); reseller keys only (403 `not_reseller`, sub-account keys 403 `subkey_forbidden`). Built for billing-panel modules (WHMCS / WISECP Gyro reseller module) that mirror the panel's lifecycle on Gyro.

API: POST /cloud/reseller/services/{id}/suspend (operationId suspendResellerService, scope reseller)

```
gyro reseller service suspend <ID> [flags]
```

### Options

```
      --data string              request body as JSON, @file or @- (stdin)
      --dry-run                  show the price quote / request without changing anything
  -h, --help                     help for suspend
      --idempotency-key string   idempotency key (automatic for money operations)
      --reason string            shown in the Gyro panel
      --set stringArray          extra body field key=value (repeatable)
```

### Options inherited from parent commands

```
      --account gyro account list   act in another account you are a panel user of (id from gyro account list; "self" = your own; env GYRO_ACCOUNT, profile default: `gyro account use`)
      --api-url string              API base URL (env GYRO_API_URL)
      --debug                       log HTTP requests to stderr (secrets masked)
      --lang string                 message language: tr or en (default from locale)
      --no-color                    disable colours (also NO_COLOR)
      --no-header                   omit the table header
  -o, --output string               output format: table, json, yaml, text
  -p, --profile string              profile to use (env GYRO_PROFILE)
      --query string                JMESPath query applied to the result
  -q, --quiet                       print only ids
      --timeout duration            overall timeout, e.g. 30s, 10m
      --wide                        show all table columns
  -y, --yes                         do not ask for confirmation
```

### SEE ALSO

* [gyro reseller service](gyro_reseller_service.md)	 - Reseller view of a service (status, due date, open invoices)

