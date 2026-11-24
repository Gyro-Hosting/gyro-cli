## gyro container egress-change approve

Approve the pending egress change

### Synopsis

Approve the pending egress change

Continue now: the new nodes start taking traffic, the old ones are drained and removed. `id` (optional) must match `egress_change.id` (409 `change_mismatch` when a newer change replaced it). 409 `no_pending_change` / `change_preparing`. Every change of a dedicated environment's node set (plan change, extra node added or removed) is announced BEFORE the old nodes go: as soon as the new nodes' addresses are known (usually under a minute) the owner is notified (Alerts channels, event `egress`) and `egress_change` / `pending_egress_ips` show them. With the approval step on (`PUT /cloud/containers/{id}/egress-approval`, 1-72 hours, 24 recommended) the new nodes take no traffic and the old nodes keep serving until `POST /cloud/containers/{id}/egress-change/approve` or the deadline (then it continues by itself). A removal only never waits.

API: POST /cloud/containers/{id}/egress-change/approve (operationId approveContainerEgressChange, scope services_write)

```
gyro container egress-change approve <ID> [flags]
```

### Options

```
      --data string              request body as JSON, @file or @- (stdin)
      --dry-run                  show the price quote / request without changing anything
  -h, --help                     help for approve
      --id string                id
      --idempotency-key string   idempotency key (automatic for money operations)
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

* [gyro container egress-change](gyro_container_egress-change.md)	 - Approve the pending egress change

