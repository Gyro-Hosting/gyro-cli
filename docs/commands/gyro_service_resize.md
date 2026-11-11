## gyro service resize

Upgrade the plan (paid from balance)

### Synopsis

Upgrade the plan (paid from balance)

dry_run returns the price; a real change needs `"confirm": true` + Idempotency-Key. Creates the upgrade invoice and pays it from the balance; if the balance does not cover it the answer is `awaiting_payment` with a pay_url. Plan changes that would re-create the server are refused (`recreate_required`). **Prepaid services** (`billing: prepaid`: Gyro Containers, Gyro Cloud Istanbul): no invoice and no confirm/Idempotency-Key; pick `product_id` (or `plan`, e.g. `standart-4-8`, `ecs-pro`) from resize-options (same product family). The module resizes the server / node pool (Istanbul servers: disks never shrink; containers: nodes of another size are replaced one by one, running services must fit a smaller plan). The new hourly price applies from the next hour; answer `status: completed`, `effective: next_hour`, `invoice: null`. A bigger plan needs 24 hours of the price difference in the balance.

API: POST /cloud/services/{id}/resize (operationId resizeService, scope cloud_order)

```
gyro service resize <ID> [flags]
```

### Options

```
      --data string              request body as JSON, @file or @- (stdin)
      --dry-run                  show the price quote / request without changing anything
  -h, --help                     help for resize
      --idempotency-key string   idempotency key (automatic for money operations)
      --plan string              prepaid services: plan name from resize-options (alternative to product_id)
      --price-id int             price_id
      --product-id int           target product from resize-options (or give plan for prepaid services)
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

