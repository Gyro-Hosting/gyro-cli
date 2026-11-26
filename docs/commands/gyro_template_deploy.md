## gyro template deploy

New server from a template

### Synopsis

New server from a template

A normal prepaid tr-ist-1 order (24-hour balance rule, account limits) installed from the template. `product_id` (or `plan`) from GET /cloud/products?region=tr-ist-1; `region` resolves the plan name (default `tr-ist-1`; Gyro Cloud Istanbul and Turkey VDS plans are both in tr-ist-1); the plan disk must be at least the template `min_disk_gb`. Send `"confirm": true` and an `Idempotency-Key` header (`dry_run` for a quote). Turkey VDS plans are billed with the plan's normal invoice: the first period (default monthly; `cycle` = a period the plan offers: monthly, semiannually, annually or biennially) is paid from the balance at order time and the order is not placed if the balance is not enough; later periods are billed as usual. `cycle` is ignored for Gyro Cloud Istanbul plans; an invalid period returns 422 `cycle_invalid` with the allowed `cycles`. On a `dry_run` quote a Turkey VDS plan returns `cycle`, `billing: "invoice"`, `total` and `tax`.

API: POST /cloud/templates/{id}/servers (operationId createServerFromTemplate, scope cloud_order)

```
gyro template deploy <ID> [flags]
```

### Options

```
      --cycle string                   Period of the first invoice, only for Turkey VDS plans: monthly (default), semiannually, annually or biennially (the periods the plan offers). Ignored for Gyro Cloud Istanbul plans; an invalid value = 422 cycle_invalid with the allowed cycles list.
      --data string                    request body as JSON, @file or @- (stdin)
      --dry-run                        show the price quote / request without changing anything
  -h, --help                           help for deploy
      --hostname string                hostname
      --idempotency-key string         idempotency key (automatic for money operations)
      --plan string                    plan
      --product-id int                 product_id
      --project-id project_not_found   Optional project for the new service(s) (0 = default project). Checked before the order: unknown id = 404 project_not_found (GyroAPI order endpoints) or 422 `checkout_blocked` with `[project_not_found]` (POST /orders).
      --region string                  Region used to resolve the plan name (default tr-ist-1; Gyro Cloud Istanbul and Turkey VDS plans are both in tr-ist-1).
      --set stringArray                extra body field key=value (repeatable)
      --ssh-keys int64Slice            ssh_keys (default [])
      --user-data string               user_data
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

* [gyro template](gyro_template.md)	 - List your templates

