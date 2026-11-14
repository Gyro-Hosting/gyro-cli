## gyro service resize

Upgrade the plan (paid from balance)

### Synopsis

Upgrade the plan (paid from balance)

dry_run returns the price; a real change needs `"confirm": true` + Idempotency-Key. Creates the upgrade invoice and pays it from the balance; if the balance does not cover it the answer is `awaiting_payment` with a pay_url. Plan changes that would re-create the server are refused (`recreate_required`). **Prepaid services** (`billing: prepaid`: Gyro Containers, Gyro Cloud Istanbul): no invoice and no confirm/Idempotency-Key; pick `product_id` (or `plan`, e.g. `standart-4-8`, `ecs-pro`) from resize-options (same product family). The module resizes the server / node pool (Istanbul servers: disks never shrink; containers: nodes of another size are replaced one by one, running services must fit a smaller plan). The new hourly price applies from the next hour; answer `status: completed`, `effective: next_hour`, `invoice: null`. A bigger plan needs 24 hours of the price difference in the balance. **Data safety (Gyro Cloud Istanbul, tr-ist-1)**: a plan change keeps the disk and all its data, the IP addresses, additional disks and snapshots. The root disk is enlarged in place and the file system is extended automatically; a disk never shrinks, so a plan with a smaller disk is refused with 409 `disk_shrink` (also in `dry_run`; resize-options marks such plans `allowed: false`). Bigger CPU/RAM is applied live when the hypervisor allows it, otherwise the server is shut down cleanly and started again (measured: 1-2 -> 2-4 restarted, about 12 seconds without SSH; 2-4 -> 4-8 live, no restart). While the server has snapshots the disk cannot grow: 409 `resize_refused` (delete the snapshots first). The answer carries `data_preserved: true`, `disk: grows_only`, `restart: if_needed`. Only a rebuild (`POST /cloud/servers/{id}/actions/rebuild`, needs `confirm: true`) erases data: the disk, additional disks and snapshots are deleted (IP addresses stay). **`recreate_required`** (422, invoiced products only): the target plan would need the service to be re-created on another platform, which deletes its data; the API never does that and nothing is changed or charged. Pick another plan from resize-options, or create a new server, move the data and delete the old one.

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

