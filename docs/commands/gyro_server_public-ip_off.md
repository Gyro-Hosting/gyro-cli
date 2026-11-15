## gyro server public-ip off

Turn the public IPv4 off

### Synopsis

Turn the public IPv4 off

Turning the public IPv4 off makes the server reachable only from its VPC, through a gateway port forward or with the web console (which always works). Confirm with the server's current IP (`confirm_ip`). Needs a ready private card (409 `no_private_network`) and no additional IPv4 (409 `extra_ips`). If the VPC gateway (NAT) is off the server loses internet access: refused with 409 `no_egress` unless `allow_no_egress: true`. The address stays reserved for the server; turning it on brings the same IP. An old guest agent answers 409 `agent_update_required` (see GET /cloud/servers/{id}/agent).

API: DELETE /cloud/servers/{id}/public-ip (operationId disableServerPublicIp, scope cloud_write)

```
gyro server public-ip off <ID> [flags]
```

### Options

```
      --allow-no-egress          allow_no_egress
      --confirm-ip string        confirm_ip
      --data string              request body as JSON, @file or @- (stdin)
      --dry-run                  show the price quote / request without changing anything
  -h, --help                     help for off
      --idempotency-key string   idempotency key (automatic for money operations)
      --set stringArray          extra body field key=value (repeatable)
      --wait                     wait until the server is no longer busy
      --wait-timeout duration    maximum time for --wait (default 15m0s)
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

* [gyro server public-ip](gyro_server_public-ip.md)	 - Public IPv4 state

