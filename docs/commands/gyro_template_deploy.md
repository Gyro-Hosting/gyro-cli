## gyro template deploy

New server from a template

### Synopsis

New server from a template

A normal prepaid tr-ist-1 order (24-hour balance rule, account limits) installed from the template. `product_id` (or `plan`) from GET /cloud/products?region=tr-ist-1; its disk must be at least the template `min_disk_gb`. Send `"confirm": true` and an `Idempotency-Key` header (`dry_run` for a quote).

API: POST /cloud/templates/{id}/servers (operationId createServerFromTemplate, scope cloud_order)

```
gyro template deploy <ID> [flags]
```

### Options

```
      --data string                    request body as JSON, @file or @- (stdin)
      --dry-run                        show the price quote / request without changing anything
  -h, --help                           help for deploy
      --hostname string                hostname
      --idempotency-key string         idempotency key (automatic for money operations)
      --plan string                    plan
      --product-id int                 product_id
      --project-id project_not_found   Optional project for the new service(s) (0 = default project). Checked before the order: unknown id = 404 project_not_found (GyroAPI order endpoints) or 422 `checkout_blocked` with `[project_not_found]` (POST /orders).
      --set stringArray                extra body field key=value (repeatable)
      --ssh-keys int64Slice            ssh_keys (default [])
      --user-data string               user_data
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

* [gyro template](gyro_template.md)	 - List your templates

