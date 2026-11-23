## gyro reseller service terminate

Terminate a service (reseller)

### Synopsis

Terminate a service (reseller)

`"confirm": true` required (422 `confirm_required`). `"immediate": true` cancels now: billing stops and the resource and its data are deleted (status `cancelling`, `cancelled` within about a minute; cannot be undone; 403 `immediate_disabled` when the operator turned it off). Without `immediate` a cancellation request is filed (active services only; a suspended service answers 409 `terminate_needs_immediate`). Cancellation rules of the product are checked first (422 `cancellation_rejected`). A cancelled service or an existing request returns `changed: false`. Only services on the reseller's own account (404 `service_not_found` otherwise); reseller keys only (403 `not_reseller`, sub-account keys 403 `subkey_forbidden`). Built for billing-panel modules (WHMCS / WISECP Gyro reseller module) that mirror the panel's lifecycle on Gyro.

API: POST /cloud/reseller/services/{id}/terminate (operationId terminateResellerService, scope reseller)

```
gyro reseller service terminate <ID> [flags]
```

### Options

```
      --data string              request body as JSON, @file or @- (stdin)
      --dry-run                  show the price quote / request without changing anything
  -h, --help                     help for terminate
      --idempotency-key string   idempotency key (automatic for money operations)
      --immediate                delete now (cannot be undone)
      --reason string            reason
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

* [gyro reseller service](gyro_reseller_service.md)	 - Reseller view of a service (status, due date, open invoices)

