## gyro reseller service get

Reseller view of a service (status, due date, open invoices)

### Synopsis

Reseller view of a service (status, due date, open invoices)

Status, product, cycle, price, due date, who suspended it, the assigned end-customer and the open invoices of the service. Only services on the reseller's own account (404 `service_not_found` otherwise); reseller keys only (403 `not_reseller`, sub-account keys 403 `subkey_forbidden`). Built for billing-panel modules (WHMCS / WISECP Gyro reseller module) that mirror the panel's lifecycle on Gyro.

API: GET /cloud/reseller/services/{id} (operationId getResellerService, scope reseller)

```
gyro reseller service get <ID> [flags]
```

### Options

```
  -h, --help   help for get
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

