## gyro container nodes set

Set the extra node count

### Synopsis

Set the extra node count

Adds or removes extra nodes of the plan's node size (absolute `extra_nodes`, 0 = plan only; at most 20 nodes per environment in total, GyroKota may lower it). An extra node costs the same as a plan node, billed hourly from the prepaid balance. Removing drains first: replicas move to the remaining nodes (services restart one by one, health checked), then the node is deleted; it is not billed while drained. Refused with 409 `quota_in_use` when the remaining nodes cannot hold the running services. Serverless environments have no nodes. Needs `"confirm": true`. Every change of a dedicated environment's node set (plan change, extra node added or removed) is announced BEFORE the old nodes go: as soon as the new nodes' addresses are known (usually under a minute) the owner is notified (Alerts channels, event `egress`) and `egress_change` / `pending_egress_ips` show them. With the approval step on (`PUT /cloud/containers/{id}/egress-approval`, 1-72 hours, 24 recommended) the new nodes take no traffic and the old nodes keep serving until `POST /cloud/containers/{id}/egress-change/approve` or the deadline (then it continues by itself). A removal only never waits.

API: PUT /cloud/containers/{id}/nodes (operationId setContainerNodes, scope services_write)

```
gyro container nodes set <ID> [flags]
```

### Options

```
      --data string              request body as JSON, @file or @- (stdin)
      --dry-run                  show the price quote / request without changing anything
      --extra-nodes int          extra_nodes
  -h, --help                     help for set
      --idempotency-key string   idempotency key (automatic for money operations)
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

* [gyro container nodes](gyro_container_nodes.md)	 - Node pool

