## gyro container egress-approval set

Approval step for egress changes

### Synopsis

Approval step for egress changes

`hours` 1-72 turns the approval step on (24 recommended), 0 turns it off. Every change of a dedicated environment's node set (plan change, extra node added or removed) is announced BEFORE the old nodes go: as soon as the new nodes' addresses are known (usually under a minute) the owner is notified (Alerts channels, event `egress`) and `egress_change` / `pending_egress_ips` show them. With the approval step on (`PUT /cloud/containers/{id}/egress-approval`, 1-72 hours, 24 recommended) the new nodes take no traffic and the old nodes keep serving until `POST /cloud/containers/{id}/egress-change/approve` or the deadline (then it continues by itself). A removal only never waits.

API: PUT /cloud/containers/{id}/egress-approval (operationId setContainerEgressApproval, scope services_write)

```
gyro container egress-approval set <ID> [flags]
```

### Options

```
      --data string              request body as JSON, @file or @- (stdin)
      --dry-run                  show the price quote / request without changing anything
  -h, --help                     help for set
      --hours int                hours
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

* [gyro container egress-approval](gyro_container_egress-approval.md)	 - Approval step for egress changes

