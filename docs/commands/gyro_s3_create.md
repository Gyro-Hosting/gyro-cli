## gyro s3 create

Order Object Storage (S3)

### Synopsis

Order Object Storage (S3)

Orders 1..N storage services of a plan in ONE order (`quantity`, default 1, max 25) and pays from the account balance (never a stored card). Billed per `cycle` (monthly | quarterly | semiannually | annually; fixed EUR prices charged in the account currency) and renewed from the balance like other services. `extra_traffic` (`1tb` | `5tb`) adds a one-time traffic pack for this month; `object_lock_buckets` (protected plans, 1-50) buys object lock rights (billed every cycle). `name` (with `{n}`), `tags` and `project_id` label the services. Send `"dry_run": true` for a quote; a real order needs `"confirm": true` and an `Idempotency-Key` header. The first bucket and the first access key are created automatically within about a minute (status `inprocess` -> `active`); then use `/cloud/storage/{id}/buckets`, `/cloud/storage/{id}/keys` and `/cloud/storage/{id}/connect`. Errors: 402 `insufficient_balance`, 409 `capacity_unavailable`, 422 `confirm_required` (with `quote`), `product_not_orderable`, `cycle_invalid`, `invalid_extra_traffic`, `invalid_object_lock`, `invalid_quantity`; 429 `daily_limit`.

API: POST /cloud/storage (operationId createStorageServices, scope cloud_order)

```
gyro s3 create [flags]
```

### Options

```
      --cycle string              billing cycle (monthly|quarterly|semiannually|annually)
      --data string               request body as JSON, @file or @- (stdin)
      --dry-run                   show the price quote / request without changing anything
      --end-customer-id int       reseller main key: assign to this end-customer
      --extra-traffic string      one-time traffic pack for this calendar month (1tb|5tb)
  -h, --help                      help for create
      --idempotency-key string    idempotency key (automatic for money operations)
      --name string               name (nickname); {n} = number in a batch
      --object-lock-buckets int   protected plans: object lock bucket rights (billed every cycle)
      --plan string               plan from GET /cloud/storage/plans
      --product-id int            product id instead of plan
      --project-id int            project (0 = default project)
      --quantity int              services to create in one order (--count), max 25
      --set stringArray           extra body field key=value (repeatable)
      --tags strings              tags
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

* [gyro s3](gyro_s3.md)	 - S3 object storage: buckets and access keys

