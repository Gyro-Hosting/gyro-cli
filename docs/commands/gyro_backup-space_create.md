## gyro backup-space create

Order a backup space

### Synopsis

Order a backup space

Orders ONE backup space and pays from the account balance (never a stored card). The space is bound to `server_id`, one of your active servers (Gyro Cloud tr-ist-1, TR VDS, managed or dedicated); it may be omitted when the account has exactly one. Only that server's IPv4 addresses can connect (SFTP, port 2201). Billed per `cycle` (fixed EUR prices charged in the account currency, VAT excluded) and renewed from the balance like other services. Send `"dry_run": true` for a quote; a real order needs `"confirm": true` and an `Idempotency-Key` header. The space is set up within about a minute (`backup_space.state` `pending` -> `active`); then read the password with `GET /cloud/backup-spaces/{id}/credentials`. Errors: 402 `insufficient_balance`; 403 `not_available` (not on sale for this account), `orders_disabled`; 409 `server_required` (no active server), `capacity_unavailable`; 422 `confirm_required` (with `quote`), `product_not_orderable`, `cycle_invalid`, `server_id_required`, `invalid_server`, `invalid_quantity`; 429 `daily_limit`.

API: POST /cloud/backup-spaces (operationId createBackupSpace, scope cloud_order)

```
gyro backup-space create [flags]
```

### Options

```
      --cycle string             billing cycle (monthly|quarterly|semiannually|annually)
      --data string              request body as JSON, @file or @- (stdin)
      --dry-run                  show the price quote / request without changing anything
      --end-customer-id int      reseller main key: assign to this end-customer
  -h, --help                     help for create
      --idempotency-key string   idempotency key (automatic for money operations)
      --name string              name (nickname)
      --plan string              plan from GET /cloud/backup-spaces/plans (50gb, 100gb, 250gb, 500gb, 1tb)
      --product-id int           product id instead of plan
      --project-id int           project (0 = default project)
      --server-id int            server to bind the space to (optional when the account has one active server)
      --set stringArray          extra body field key=value (repeatable)
      --tags strings             tags
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

* [gyro backup-space](gyro_backup-space.md)	 - Backup Space: SFTP backup space bound to one of your servers (tr-ist-1)

